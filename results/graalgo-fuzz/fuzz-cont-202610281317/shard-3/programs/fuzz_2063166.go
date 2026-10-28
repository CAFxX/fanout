package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint8, p2 uint) int {
    v7 := []int{-966, 127}
    p1 -= 255
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 9223372036854775807
}
func genFn1() int16 {
    v8 := []int{348}
    if false {
        var v9 float32 = 0.0
        var v10 int16 = -844
        _ = v10
        _ = v9
    } else {
        var v11 int32 = 2147483647
        v11 += -1
        _ = v11
    }
    var v12 float32 = ((3.14159 - (-2.25 / 1.5)) + 1.5)
    _ = v12
    _ = v8
    return -1
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]uint32)
    var v3 int8 = 42
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    v6 := []string{"\n\t", ("" + "hello")}
    v3 += (v3 + v3)
    for fi13 := 0; fi13 < 5; fi13++ {
        if (220 != 185) {
            fmt.Println(v1, v3, fi13)
        } else {
            switch "a" {
            case "foo bar":
                v14 := GenS1{F0: 9223372036854775807}
                _ = v14
            case "a":
                for fi15 := 0; fi15 < 19; fi15++ {
                    if false {
                        fmt.Println(fi13, fi15)
                        var v16 uint8 = 255
                        _ = v16
                    }
                    var v17 uint64 = 0
                    delete(v4, 4294967295)
                    _ = fi15
                    _ = v17
                }
            case "":
                switch "foo bar" {
                case "\n\t":
                    var v18 int = (fi13 >> ((65535) & 63))
                    _ = v18
                case "hello":
                    var v19 uint64 = 1
                    _ = v19
                }
            default:
                if ((561 > 883) || false) {
                    fmt.Println(fi13, fi13)
                }
            }
            fmt.Println(fi13)
        }
        fmt.Println(fi13, v1)
        _ = fi13
    }
    switch ("" + ("hello" + "\n\t")) {
    case "a":
        switch 4294967295 {
        case -9223372036854775808:
            v1 -= -2.25
        default:
            _ = len(v5)
        }
    case "\n\t":
        if true {
            var v20 int64 = 42
            var v21 int64 = (v20 + (v20 / ((v20) | 1)))
            var v22 int64 = v21
            _ = v20
            _ = v21
            _ = v22
        }
    }
    v1 = v1
    for fi23 := 0; fi23 < 18; fi23++ {
        if ((781 > 1000) && false) {
            var v24 int8 = v3
            switch "hello" {
            case "hello":
                if (true || (-743 < -2147483648)) {
                    _ = len(v6)
                }
            case "":
                v3 = v24
            case "a":
                fmt.Println(v24, fi23)
            }
            v25 := GenS1{F0: 9223372036854775807}
            _ = v24
            _ = v25
        }
        _ = len(v6)
        _ = fi23
    }
    v26 := GenS1{F0: -9223372036854775808}
    _ = len(v6)
    v3 -= (v3 + v3)
    if false {
        switch 127 {
        case 9223372036854775807:
            v1 -= v1
        }
        if (!(v3 >= v3)) {
            fmt.Println(v1)
            v3 = 127
        }
        fmt.Println(v3, v1)
    }
    var v27 int64 = int64(v3)
    v1 = (-2.25 / (v1 / (0.0 - 1.5)))
    fmt.Println(v27)
    if len(v6) > 0 {
        v6[0] = "foo bar"
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
