package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint16; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int }

func main() {
    var v1 float64 = 1.5
    var v2 int64 = -9223372036854775808
    v3 := []string{}
    var v4 bool = true
    v5 := GenS1{F0: 111, F1: 161}
    v6 := GenS1{F0: -110, F1: 2}
    v7 := GenS1{F0: -1, F1: 1}
    v1 = 1.5
    _ = len(v3)
    v8 := GenS2{F0: v4, F1: 255, F2: ("a" + "")}
    if ((false && v4) && (!v4)) {
        v9 := GenS1{F0: 0, F1: 93}
        switch ("hello" + "foo bar") {
        case "":
            var v10 string = "hello"
            _ = v10
        case "\n\t":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "hello":
            fmt.Println(v2)
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            fmt.Println(v4, fi11)
            var v12 float32 = 0.0
            _ = fi11
            _ = v12
        }
        _ = v9
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS3{F0: 0}
        var v15 string = ""
        _ = fi13
        _ = v14
        _ = v15
    }
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v16 float64 = v1
        _ = v16
    case 0:
        v17 := GenS2{F0: (-446 < 9223372036854775807), F1: 76, F2: "foo bar"}
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 uint16 = 255
            fmt.Println(v4)
            if (((!(!((false || (-899 > 32767)) && false))) || ((v2 <= 1) || (!false))) && (171 > 2)) {
                v20 = 65535
                if (1000 <= 42) {
                    v21 := GenS1{F0: -128, F1: 230}
                    _ = v21
                }
            } else {
                var v22 uint32 = 344
                v23 := GenS4{F0: (v20 << ((uint(v22)) & 63)), F1: fi19}
                _ = v22
                _ = v23
            }
            _ = fi19
            _ = v20
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 uint16 = 259
            _ = fi24
            _ = v25
        }
        _ = len(v3)
        _ = fi18
    }
    v3 = append(v3, "hello")
    switch "\n\t" {
    case "hello":
        var v26 int32 = 1
        _ = v26
    case "foo bar":
        v4 = (!((("hello" + "\n\t") + ("hello" + "foo bar")) <= "foo bar"))
    default:
        v2 += v2
    }
    if (32767 >= 365) {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
