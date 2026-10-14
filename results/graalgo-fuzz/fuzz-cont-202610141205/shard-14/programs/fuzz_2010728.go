package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint) bool {
    switch (((("hello" + ((("hello" + ("hello" + ("a" + ("\n\t" + (("foo bar" + ("" + (((("\n\t" + (("\n\t" + "a") + "\n\t")) + "foo bar") + ((("" + ("\n\t" + "hello")) + "a") + ("hello" + ((("\n\t" + "a") + ("hello" + ((("a" + "") + "foo bar") + "a"))) + ("" + ("hello" + "")))))) + ("a" + "")))) + "foo bar"))))) + "\n\t") + "a")) + ("foo bar" + (("a" + ("a" + "\n\t")) + "hello"))) + ("hello" + ("" + "a"))) + "\n\t") {
    case "\n\t":
        v4 := GenS1{F0: 57}
        _ = v4
    default:
        p0 = p0
    }
    p0 += 2
    fmt.Println(p0)
    _ = p0
    return true
}

func main() {
    var v1 int8 = -82
    var v2 string = "hello"
    var v3 float64 = -2.25
    v5 := GenS1{F0: -1}
    if (2147483647 <= 2147483647) {
        if (!true) {
            v2 = v2
        } else {
            v6 := GenS1{F0: 2147483647}
            v3 -= 0.1
            fmt.Println(v1)
            _ = v6
        }
    } else {
        v7 := []int{-1, 480, 9223372036854775807, 32767}
        switch "hello" {
        case "":
            var v8 uint16 = 1
            _ = v8
        default:
            var v9 uint32 = 2
            _ = v9
        }
        _ = v7
    }
    v2 = "a"
    if true {
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 float32 = 1.5
            _ = fi10
            _ = v11
        }
        v12 := []int{}
        fmt.Println(v3, v2, v2)
        _ = v12
    } else {
        var v13 int32 = 42
        v14 := []int{-9223372036854775808, 2147483647, 9223372036854775807, 1}
        var v15 float64 = 1.5
        _ = v13
        _ = v14
        _ = v15
    }
    v1 = (v1 / (((v1 / (((v1 ^ v1)) | 1))) | 1))
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(v1, fi16)
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 float64 = (v3 * v3)
            fmt.Println(v1)
            if (47 != 1) {
                v19 := []int{(fi17 / ((fi16) | 1)), (^(fi17)), (^(fi17)), (fi16 - (fi17 / ((fi17) | 1)))}
                fmt.Println(v2, v2, v2)
                v20 := GenS1{F0: -2147483648}
                _ = v19
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v1 = -128
    v21 := []int{}
    switch v2 {
    case "hello":
        if true {
            var v22 uint16 = 345
            for fi23 := 0; fi23 < 7; fi23++ {
                if ((!(true && true)) || ((0.0 - 3.14159) <= 3.14159)) {
                    v21 = append(v21, 0)
                } else {
                    v21 = append(v21, (fi23 / ((936) | 1)))
                }
                _ = fi23
            }
            _ = v22
        } else {
            fmt.Println(v3)
            var v24 float32 = (1e10 + 0.0)
            for fi25 := 0; fi25 < 20; fi25++ {
                v24 += ((v24 / (v24 / 0.1)) * (0.0 * float32(v3)))
                v2 = v2
                _ = fi25
            }
            _ = v24
        }
    case "foo bar":
        v2 = "hello"
    default:
        v26 := GenS1{F0: 0}
        _ = v26
    }
    var v27 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v5
}
