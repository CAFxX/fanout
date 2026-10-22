package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int32; F2 uint64 }

func main() {
    v1 := []int{431, -450, 525, -2147483648}
    v2 := make(map[uint32]int)
    var v3 uint16 = 999
    var v4 int64 = -1
    var v5 int8 = 106
    var v6 uint = 1
    v7 := GenS1{F0: 42, F1: 128}
    fmt.Println(v4, v6)
    var v8 int32 = -1
    v8 += (-(v8))
    var v9 bool = (false && true)
    var v10 string = ((("foo bar" + "foo bar") + ("foo bar" + ("\n\t" + "\n\t"))) + "")
    if (v9 && v9) {
        if (!v9) {
            var v11 int8 = (v5 + v5)
            if (!(127 != 42)) {
                v1 = append(v1, 4294967295)
            } else {
                if true {
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                }
            }
            v10 = v10
            _ = v11
        } else {
            if (((v6 + (^(v6))) < (v6 << (((v6 - (^(v6)))) & 63))) && (!(9223372036854775807 == 490))) {
                v12 := GenS1{F0: 1, F1: -232}
                fmt.Println(v10)
                _ = v12
            } else {
                fmt.Println(v6, v3, v6)
                v3 = (v3 % (((v3 / ((v3) | 1))) | 1))
            }
            var v13 int16 = 1000
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v13
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            v9 = false
            _ = fi14
        }
        var v15 int8 = -128
        _ = v15
    } else {
        switch v10 {
        case "hello":
            var v16 uint16 = (v3 | 116)
            _ = v16
        case "a":
            v17 := GenS2{F0: -466}
            _ = v17
        case "":
            var v18 float64 = (1.5 + (3.14159 / (3.14159 / 1.5)))
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        v6 = 255
        _ = fi19
    }
    var v20 string = ("foo bar" + ("hello" + "foo bar"))
    var v21 string = v10
    v22 := GenS3{F0: (v8 + v8), F1: (v8 - (v8 * v8)), F2: 721}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
