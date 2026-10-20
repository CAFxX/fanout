package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 string }
func genFn0(p0 int32, p1 bool) int32 {
    if true {
        p0 = p0
        v7 := []int{32767, 612, 1}
        var v8 uint16 = 65535
        _ = v7
        _ = v8
    } else {
        v9 := []int{}
        var v10 uint8 = 127
        _ = v10
        _ = v9
    }
    var v11 int16 = -32768
    fmt.Println(p0, p1)
    fmt.Println(v11)
    _ = p0
    _ = p1
    _ = v11
    return 2147483647
}

func main() {
    var v1 float64 = float64(-2.25)
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    var v4 string = ""
    var v5 uint32 = 4294967295
    v6 := make(map[uint32]uint32)
    v5 += v5
    v4 = v4
    switch -2147483648 {
    case -33:
        var v12 float32 = 1.5
        _ = v12
    case -1:
        var v13 int8 = -1
        _ = v13
    default:
        v14 := GenS1{F0: 0}
        _ = v14
    }
    v3[(^(v5))] = false
    v4 = ((("hello" + (("hello" + ((v4 + "") + "a")) + "")) + v4) + v4)
    var v15 int16 = 1000
    if false {
        var v16 float64 = 0.0
        _ = v16
    } else {
        v4 = v4
        switch -2147483648 {
        case 42:
            if (true || (false || (55 == -1))) {
                var v17 uint32 = 1
                _ = v17
            }
        case -652:
            fmt.Println(v15, v1)
        default:
            var v18 uint16 = 65535
            _ = v18
        }
        switch "foo bar" {
        case "\n\t":
            if (42 <= 18446744073709551615) {
                if (!(1 >= genFn0(-2147483648, true))) {
                    var v19 int16 = v15
                    switch v4 {
                    case "\n\t":
                        v20 := GenS1{F0: 4294967295}
                        _ = v20
                    default:
                        v21 := GenS1{F0: 1}
                        _ = v21
                    }
                    _ = v19
                }
                v4 = v4
                v15 += v15
            } else {
                for fi22 := 0; fi22 < 4; fi22++ {
                    var v23 int32 = -1
                    v24 := GenS1{F0: 1}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
        default:
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 int16 = (^(v15))
                if true {
                    for fi27 := 0; fi27 < 5; fi27++ {
                        var v28 int16 = (v15 & (v15 / ((-43) | 1)))
                        var v29 bool = ((v5 - (^(v5))) != (v5 ^ (v5 >> ((487) & 63))))
                        var v30 string = ("a" + "")
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    var v31 int8 = int8(v15)
                    _ = v31
                } else {
                    v32 := GenS1{F0: 65535}
                    if true {
                        v33 := GenS1{F0: 965}
                        _ = v33
                    }
                    v34 := GenS2{F0: -9223372036854775808, F1: 1}
                    _ = v32
                    _ = v34
                }
                v35 := GenS3{F0: ""}
                _ = fi25
                _ = v26
                _ = v35
            }
        }
    }
    v4 = (v4 + v4)
    var v36 int32 = 42
    var v37 int = 641
    fmt.Println(v37, v15, v15)
    var v38 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
