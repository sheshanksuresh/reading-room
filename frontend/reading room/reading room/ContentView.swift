//
//  ContentView.swift
//  reading room
//
//  Created by Sheshank Personal on 2024-02-10.
//

import SwiftUI

struct ContentView: View {
    @State private var username: String = ""
    @State private var password: String = ""
    
    var body: some View {
        NavigationView(content: {
            VStack(content: {
                Text("Welcome to the Reading Room!")
                    .font(.largeTitle)
                    .padding()
                
                TextField("Username", text: $username)
                    .textFieldStyle(.roundedBorder)
                    .padding()
                
                TextField("Password", text: $password)
                    .textFieldStyle(.roundedBorder)
                    .padding()
                
                Button("Login", action: {})
                    .padding()
                    .font(.title3)
                    .fontWeight(.heavy)
                    .frame(width: 250)
                    .background(Color.blue)
                    .foregroundColor(.white)
                    .cornerRadius(40)
                
                Spacer().frame(height: 20)
                
                HStack(content: {
                    Text("Don't have an account?")
                    
                    NavigationLink(destination: CreateAccountView()) {
                        Text("Create Account")
                    }
                })
                
                
                
            })
        })
        
    }
}

#Preview {
    ContentView()
}
