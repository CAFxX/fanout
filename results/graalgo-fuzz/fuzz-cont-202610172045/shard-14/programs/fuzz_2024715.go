package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int16) bool {
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int32 = 0
        if (0 != 65535) {
            if ((fi4 & 0) != fi4) {
                v6 := GenS1{F0: "\n\t"}
                v7 := []int{(-(fi4)), -9223372036854775808}
                _ = v6
                _ = v7
            }
            v5 -= (v5 >> ((65535) & 63))
        }
        _ = fi4
        _ = v5
    }
    v8 := []int{}
    _ = p0
    _ = v8
    return ((1 > 2147483647) && (!(false && false)))
}

func main() {
    var v1 uint16 = 32767
    var v2 int32 = -802
    var v3 uint = 1
    v3 += (v3 >> ((v3) & 63))
    fmt.Println(v3, v3)
    if (!(false && true)) {
        var v9 int = 1
        if ((^(v1)) > 65535) {
            fmt.Println(v9)
            v10 := GenS1{F0: "foo bar"}
            fmt.Println(v1, v9)
            _ = v10
        }
        v9 += v9
        _ = v9
    } else {
        v11 := GenS1{F0: "\n\t"}
        _ = v11
    }
    var v12 int8 = 42
    for fi13 := 0; fi13 < 20; fi13++ {
        v1 = 42
        if true {
            v14 := GenS1{F0: "foo bar"}
            var v15 uint64 = 1
            _ = v14
            _ = v15
        }
        v16 := []int{}
        _ = fi13
        _ = v16
    }
    v17 := GenS2{F0: (v12 ^ 42)}
    for fi18 := 0; fi18 < 10; fi18++ {
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v2, v1, fi18)
        case "hello":
            var v19 uint16 = 255
            _ = v19
        default:
            for fi20 := 0; fi20 < 3; fi20++ {
                v21 := GenS3{F0: (^(v12))}
                v1 = v1
                _ = fi20
                _ = v21
            }
        }
        switch "\n\t" {
        case "\n\t":
            var v22 string = "hello"
            _ = v22
        case "":
            for fi23 := 0; fi23 < 13; fi23++ {
                var v24 int32 = (v2 + v2)
                v25 := []int{691, (fi18 + 32767), (fi18 & 2147483647)}
                _ = fi23
                _ = v24
                _ = v25
            }
        default:
            var v26 int32 = 244
            _ = v26
        }
        _ = fi18
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        fmt.Println(v12, fi27, v1)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
}
