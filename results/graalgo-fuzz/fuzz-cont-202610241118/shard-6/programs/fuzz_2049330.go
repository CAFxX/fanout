package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int8; F1 int16; F2 int8 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]string)
    var v3 string = "hello"
    var v4 uint16 = 65535
    var v5 float32 = v1
    var v6 uint16 = 42
    if false {
        var v7 int64 = -793
        _ = v7
    }
    v8 := GenS1{F0: 4294967295, F1: -1}
    v6 = (^(v4))
    _ = len(v2)
    v4 = (^(v4))
    var v9 uint32 = 688
    if true {
        if (443 > 1) {
            if true {
                var v10 int8 = -36
                _ = v10
            }
        } else {
            var v11 string = ((v3 + v3) + v3)
            _ = v11
        }
    }
    v4 += 32767
    switch (((v3 + ("\n\t" + (v3 + v3))) + ("a" + ((((v3 + v3) + ("\n\t" + "\n\t")) + v3) + v3))) + (v3 + ((v3 + "a") + (v3 + "a")))) {
    case "a":
        var v12 uint16 = 1
        _ = v12
    }
    if (v3 <= v3) {
        v6 -= 0
    } else {
        v2[((v3 + "") + (v3 + v3))] = v3
        var v13 uint8 = 42
        _ = v13
    }
    if (true || (-685 < -2147483648)) {
        v14 := GenS1{F0: v9, F1: -128}
        v15 := GenS2{F0: -128, F1: -692, F2: 1}
        _ = len(v2)
        _ = v14
        _ = v15
    }
    v16 := v2[v3]
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
