package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 int; F2 string }

func main() {
    var v1 uint64 = 4294967295
    var v2 int32 = -1
    v3 := make(map[int]uint32)
    var v4 uint8 = 162
    var v5 uint = 1
    for fi6 := 0; fi6 < 18; fi6++ {
        v4 -= v4
        if (42 <= 690) {
            v1 = (^(v1))
        } else {
            switch (("a" + ("hello" + "a")) + "a") {
            case "foo bar":
                v5 = 1
            case "hello":
                var v7 float32 = -2.25
                _ = v7
            case "\n\t":
                v8 := GenS1{F0: 516}
                _ = v8
            }
            v9 := GenS1{F0: 1}
            v10 := GenS2{F0: (^(v1))}
            _ = v10
            _ = v9
        }
        switch "\n\t" {
        case "foo bar":
            v5 = v5
        }
        _ = fi6
    }
    var v11 float32 = float32(1.5)
    fmt.Println(v2, v5)
    v12 := v3[-595]
    v13 := GenS1{F0: v12}
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := v3[int(v2)]
            var v17 uint16 = 1
            v18 := []int{32768, fi14, (^(fi15))}
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi14
    }
    v19 := GenS2{F0: v1}
    fmt.Println(v1, v4)
    if true {
        fmt.Println(v11)
        if false {
            fmt.Println(v11)
        } else {
            if true {
                if ((!(false && false)) || false) {
                    v20 := GenS3{F0: 1}
                    var v21 uint16 = 32767
                    _ = v20
                    _ = v21
                }
            }
            v22 := GenS4{F0: (("foo bar" + "a") + ((("a" + "hello") + "") + ("hello" + "hello")))}
            var v23 int = 0
            _ = v22
            _ = v23
        }
        switch -197 {
        case 32768:
            v24 := GenS4{F0: ""}
            _ = v24
        case 32767:
            v25 := GenS5{F0: (v2 | v2), F1: 9223372036854775807, F2: ("hello" + "hello")}
            _ = v25
        default:
            fmt.Println(v11)
        }
    }
    v26 := GenS1{F0: (v12 ^ (v12 % (((v12 ^ 2147483647)) | 1)))}
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
