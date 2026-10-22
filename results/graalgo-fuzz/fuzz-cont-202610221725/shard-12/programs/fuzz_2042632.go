package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
func genFn0(p0 string, p1 uint32, p2 int64) bool {
    v5 := GenS1{F0: "a"}
    if (((1e10 + 0.0) * 1e10) > 0.0) {
        if true {
            for fi6 := 0; fi6 < 13; fi6++ {
                p1 -= 2147483647
                var v7 uint = 0
                _ = fi6
                _ = v7
            }
            fmt.Println(p2, p0)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return false
}

func main() {
    var v1 uint8 = 223
    var v2 bool = true
    v3 := []int{-1}
    var v4 int64 = 42
    v8 := GenS1{F0: "a"}
    v9 := GenS1{F0: (("a" + "") + "a")}
    _ = len(v3)
    var v10 uint = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v4, v2)
    switch "hello" {
    case "hello":
        v10 = v10
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "":
        switch 32767 {
        case -2147483648:
            switch -9223372036854775808 {
            case -1:
                var v11 int32 = 1
                _ = v11
            default:
                v10 -= (v10 | (v10 / (((v10 % (((v10 + 4294967295)) | 1))) | 1)))
            }
        case 9223372036854775807:
            v1 += 0
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v10)
    if true {
        switch ("" + "hello") {
        case "":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "hello":
            switch ("hello" + ((("" + "\n\t") + "hello") + "a")) {
            case "foo bar":
                v10 += 42
            case "hello":
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 int64 = int64(v1)
                    _ = fi12
                    _ = v13
                }
            case "a":
                fmt.Println(v4, v1, v2)
            }
        case "a":
            if v2 {
                var v14 uint8 = 115
                _ = v14
            }
        }
        if (-961 > -1) {
            v3 = append(v3, 9223372036854775807)
            v10 = (v10 | uint(v1))
        }
    } else {
        if len(v3) > 0 {
            v3[0] = -728
        }
        if v2 {
            v2 = (true && (!v2))
            var v15 float64 = float64(3.14159)
            _ = v15
        } else {
            if len(v3) > 0 {
                v3[0] = 42
            }
            v16 := GenS1{F0: "a"}
            _ = v16
        }
        var v17 int = 337
        _ = v17
    }
    if (-1 == 0) {
        v4 = v4
        v18 := GenS2{F0: 208}
        _ = v18
    }
    v1 = (v1 & v1)
    var v19 uint64 = 246
    if false {
        switch -1 {
        case 9223372036854775807:
            var v20 uint64 = (v19 | v19)
            _ = v20
        case 2147483647:
            fmt.Println(v4)
        default:
            _ = len(v3)
        }
    }
    var v21 int32 = -259
    v22 := GenS3{F0: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
