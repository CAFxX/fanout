package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 bool = (true || false)
    var v2 int32 = 1
    var v3 string = ("foo bar" + ("" + "hello"))
    var v4 bool = v1
    var v5 string = v3
    v6 := []string{(v3 + v3), ("" + "\n\t"), "hello"}
    var v7 int = -344
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v5 = ((((v5 + (v5 + v3)) + "a") + ("a" + v3)) + v3)
    v8 := GenS1{F0: 42}
    v4 = true
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 float32 = 0.1
        _ = len(v6)
        _ = fi9
        _ = v10
    }
    v3 = (("\n\t" + v3) + v5)
    for fi11 := 0; fi11 < 2; fi11++ {
        if len(v6) > 0 {
            v6[0] = ((v5 + (v3 + ("a" + ((v5 + (v3 + "\n\t")) + v5)))) + v5)
        }
        if false {
            v3 = "a"
            v12 := GenS1{F0: 1}
            _ = v12
        } else {
            switch fi11 {
            case 127:
                var v13 int = v7
                _ = v13
            case 894:
                var v14 uint64 = 315
                _ = v14
            case -9223372036854775808:
                v3 = (v5 + v3)
            }
        }
        var v15 bool = (!v4)
        _ = fi11
        _ = v15
    }
    v16 := GenS1{F0: 2}
    v3 = v5
    v1 = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
