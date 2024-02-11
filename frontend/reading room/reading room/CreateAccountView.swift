//
//  CreateAccountView.swift
//  Reading Room
//
//  Created by Sheshank Personal on 2024-02-10.
//

import SwiftUI

struct CreateAccountView: View {
    @State private var firstName: String = ""
    @State private var lastName: String = ""
    @State private var username: String = ""
    @State private var password: String = ""
    @State private var email: String = ""
    
    var body: some View {
        ScrollView {
            TextField("First Name", text: $firstName)
                .textFieldStyle(.roundedBorder)
                .padding()
            
            TextField("Last Name", text: $lastName)
                .textFieldStyle(.roundedBorder)
                .padding()
            
            TextField("Username", text: $username)
                .textFieldStyle(.roundedBorder)
                .padding()
            
            SecureField("Password", text: $password)
                .textFieldStyle(.roundedBorder)
                .padding()
            
            TextField("Email", text: $email)
                .textFieldStyle(.roundedBorder)
                .padding()
            
             Button("Create Account", action: {
                 
             })
                 .padding()
                 .font(.title3)
                 .fontWeight(.heavy)
                 .frame(minWidth: 0, maxWidth: .infinity)
                 .background(Color.blue)
                 .foregroundColor(.white)
                 .cornerRadius(40)
        }
    }
}

#Preview {
    CreateAccountView()
}
