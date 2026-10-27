package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int, p1 int32, p2 uint32) int64 {
    v5 := GenS1{F0: 0}
    for fi6 := 0; fi6 < 6; fi6++ {
        p1 -= (p1 / ((p1) | 1))
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{128, fi7, fi7, fi7}
            _ = fi7
            _ = v8
        }
        p1 -= (p1 & p1)
        _ = fi6
    }
    var v9 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return 1
}
func genFn1() uint32 {
    var v10 int32 = 42
    if true {
        for fi11 := 0; fi11 < 2; fi11++ {
            switch "hello" {
            case "":
                var v12 float64 = float64(0.1)
                _ = v12
            default:
                var v13 bool = false
                _ = v13
            }
            _ = fi11
        }
        var v14 bool = true
        var v15 uint64 = 18446744073709551615
        _ = v14
        _ = v15
    } else {
        switch (("foo bar" + ("hello" + "a")) + "\n\t") {
        case "":
            v16 := GenS1{F0: 0}
            _ = v16
        case "a":
            if true {
                for fi17 := 0; fi17 < 14; fi17++ {
                    v10 = v10
                    _ = fi17
                }
                if false {
                    v10 = v10
                }
                switch 2147483647 {
                case -1:
                    fmt.Println(v10)
                }
            } else {
                v18 := GenS1{F0: 437}
                var v19 int = -1
                v20 := GenS1{F0: 0}
                _ = v18
                _ = v19
                _ = v20
            }
        case "\n\t":
            if (("foo bar" + ("hello" + ("hello" + ((("" + "hello") + "foo bar") + "")))) <= ("hello" + "\n\t")) {
                v10 = (v10 + (v10 % ((v10) | 1)))
                fmt.Println(v10)
            } else {
                v21 := []int{32768, 128, 902}
                _ = v21
            }
        }
    }
    switch "\n\t" {
    case "":
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 int64 = -9223372036854775808
            _ = fi22
            _ = v23
        }
    }
    v24 := []int{161, 2147483647, 9223372036854775807, -2147483648}
    v10 = v10
    _ = v10
    _ = v24
    return 351
}

func main() {
    var v1 uint = 42
    v2 := []int{849, 9223372036854775807, -1, 9223372036854775807}
    var v3 float32 = -2.25
    v4 := make(map[uint32]bool)
    v3 -= v3
    switch "hello" {
    case "\n\t":
        v1 = (v1 / (((v1 - 0)) | 1))
    case "hello":
        var v25 int64 = 42
        _ = v25
    }
    if true {
        var v26 bool = false
        _ = v26
    } else {
        v3 += ((v3 / ((1e10 / v3) * v3)) - float32(3.14159))
        var v27 int32 = 2147483647
        _ = v27
    }
    var v28 uint = (v1 * (v1 + 18446744073709551615))
    var v29 int8 = 0
    for fi30 := 0; fi30 < 8; fi30++ {
        v3 -= (v3 / v3)
        var v31 int64 = -534
        _ = fi30
        _ = v31
    }
    var v32 int16 = 32767
    var v33 string = "a"
    v2 = append(v2, 629)
    v34 := v4[383]
    var v35 uint64 = 18446744073709551615
    var v36 uint64 = v35
    fmt.Println(v1)
    for fi37 := 0; fi37 < 4; fi37++ {
        switch 128 {
        case -324:
            var v38 float32 = ((v3 * v3) - 0.0)
            _ = v38
        default:
            v28 = (v1 << ((uint(v36)) & 63))
        }
        v35 -= (v36 & v36)
        fmt.Println(v33)
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
