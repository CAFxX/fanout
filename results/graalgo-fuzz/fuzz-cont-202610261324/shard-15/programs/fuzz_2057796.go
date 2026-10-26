package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 32767
    var v3 bool = (127 != 196)
    switch "foo bar" {
    case "hello":
        for fi4 := 0; fi4 < 9; fi4++ {
            var v5 int16 = v2
            var v6 int = 1
            _ = fi4
            _ = v5
            _ = v6
        }
    case "a":
        switch ("hello" + "foo bar") {
        case "a":
            v7 := []int{709}
            _ = v7
        }
    case "\n\t":
        v8 := []int{1}
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: 371}
    v3 = ((356 > 4294967295) || ((v3 || ((0.0 - 1.5) <= (0.0 - -2.25))) && (-2.25 < 3.14159)))
    for fi10 := 0; fi10 < 5; fi10++ {
        fmt.Println(v2)
        switch (fi10 | (fi10 | (fi10 >> ((65535) & 63)))) {
        case -949:
            switch fi10 {
            case 1:
                v2 += (v2 | v2)
            case 98:
                v11 := []int{}
                _ = v11
            default:
                if (108 < 0) {
                    var v12 int8 = int8(v2)
                    v13 := []int{}
                    _ = v12
                    _ = v13
                } else {
                    if ((fi10 - (fi10 >> ((100) & 63))) >= -9223372036854775808) {
                        v14 := GenS1{F0: 16, F1: 65535}
                        v15 := []int{(^(fi10)), 128}
                        _ = v14
                        _ = v15
                    } else {
                        var v16 int32 = 1
                        _ = v16
                    }
                    v3 = v3
                }
            }
        case 42:
            for fi17 := 0; fi17 < 5; fi17++ {
                fmt.Println(v2, fi10)
                _ = fi17
            }
        case 32768:
            if true {
                v18 := []int{fi10, 128}
                _ = v18
            } else {
                v2 -= -32768
            }
        }
        var v19 bool = ((v2 % ((v2) | 1)) > -32768)
        _ = fi10
        _ = v19
    }
    var v20 string = "foo bar"
    fmt.Println(v3, v2, v20)
    fmt.Println(v20)
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 uint32 = 2147483647
        _ = fi21
        _ = v22
    }
    v2 = 1
    var v23 int = 128
    switch 128 {
    case 42:
        var v24 float32 = 3.14159
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v9
}
