package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int32; F2 string }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint8) uint16 {
    v4 := GenS1{F0: -2147483648}
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: 0}
        _ = fi5
        _ = v6
    }
    switch 4294967295 {
    case 127:
        var v7 uint16 = 65535
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = -457
    var v3 uint16 = 32767
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int32 = -1
        _ = fi8
        _ = v9
    }
    fmt.Println(v3)
    v1[2] = 0
    if false {
        v3 = v3
    }
    if (42 == 2147483648) {
        v10 := GenS1{F0: -1}
        _ = v10
    }
    var v11 uint64 = uint64(v3)
    if false {
        for fi12 := 0; fi12 < 3; fi12++ {
            v3 -= 32767
            fmt.Println(fi12, fi12)
            _ = fi12
        }
        v13 := GenS2{F0: 56}
        _ = v13
    } else {
        v3 = 2
        fmt.Println(v11)
    }
    v14 := GenS2{F0: -20}
    fmt.Println(v11, v11, v3)
    switch "" {
    case "a":
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 uint = 4294967295
            v17 := GenS1{F0: int32(v3)}
            var v18 bool = false
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
    case "\n\t":
        for fi19 := 0; fi19 < 3; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := GenS2{F0: 32767}
                if true {
                    var v22 int8 = 1
                    var v23 string = ("hello" + ("a" + ""))
                    _ = v22
                    _ = v23
                } else {
                    var v24 string = (("a" + "foo bar") + ("foo bar" + ("foo bar" + (("" + ("hello" + "a")) + "foo bar"))))
                    _ = v24
                }
                v25 := []int{fi20}
                _ = fi20
                _ = v21
                _ = v25
            }
            v26 := GenS3{F0: 2, F1: 42, F2: "foo bar"}
            _ = fi19
            _ = v26
        }
    default:
        for fi27 := 0; fi27 < 15; fi27++ {
            fmt.Println(fi27, v11, v3)
            if false {
                var v28 uint16 = 2
                _ = v28
            } else {
                var v29 uint16 = v3
                if (1.5 < ((float64(-2.25) / (float64(3.14159) - 1.5)) + -2.25)) {
                    delete(v1, 2147483647)
                    v30 := GenS2{F0: fi27}
                    _ = v30
                }
                var v31 int8 = -128
                _ = v29
                _ = v31
            }
            _ = fi27
        }
    }
    fmt.Println(v3)
    v32 := v1[65535]
    if false {
        var v33 int64 = 1
        fmt.Println(v32, v3, v32)
        for fi34 := 0; fi34 < 9; fi34++ {
            v1[65535] = 32767
            v3 += (v3 & (^(v3)))
            v33 -= v33
            _ = fi34
        }
        _ = v33
    } else {
        v35 := GenS4{F0: (^(v11))}
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v32
}
