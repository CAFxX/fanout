package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 int32; F2 uint8 }

func main() {
    var v1 uint32 = 901
    v2 := make(map[int]int)
    var v3 string = "hello"
    var v4 float64 = -2.25
    v5 := GenS1{F0: ("\n\t" + ((v3 + "hello") + v3))}
    fmt.Println(v1, v4)
    if true {
        if true {
            v1 = 949
            v6 := GenS2{F0: ""}
            _ = v6
        } else {
            v7 := v2[-9223372036854775808]
            _ = v7
        }
    } else {
        v8 := GenS3{F0: (v1 >> ((65535) & 63))}
        v9 := GenS2{F0: v3}
        v3 = (("" + (("hello" + v3) + v3)) + v3)
        _ = v8
        _ = v9
    }
    v10 := GenS4{F0: (v3 > "\n\t"), F1: 548, F2: 1}
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 uint64 = 945
        if false {
            for fi13 := 0; fi13 < 7; fi13++ {
                switch "foo bar" {
                case "hello":
                    v3 = v3
                default:
                    fmt.Println(v4)
                }
                var v14 uint64 = (v12 >> ((42) & 63))
                v3 = "a"
                _ = fi13
                _ = v14
            }
            v2[(fi11 / ((fi11) | 1))] = fi11
            var v15 int64 = int64(fi11)
            _ = v15
        } else {
            if (false || false) {
                v16 := []int{fi11, fi11, -1}
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 float32 = (0.1 + 1e10)
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            v4 = 1e10
            v19 := []int{}
            _ = v19
        }
        _ = fi11
        _ = v12
    }
    v1 = (v1 ^ v1)
    var v20 int8 = -128
    for fi21 := 0; fi21 < 4; fi21++ {
        fmt.Println(v20, v1)
        var v22 bool = true
        v23 := []int{(fi21 / ((32768) | 1)), (fi21 ^ fi21)}
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v3, v20)
    var v24 int = 1
    v24 -= -722
    var v25 uint64 = 0
    v3 = (v3 + v3)
    if (!(255 != 127)) {
        var v26 int32 = 42
        _ = v26
    }
    _ = len(v2)
    if true {
        for fi27 := 0; fi27 < 8; fi27++ {
            if ((v20 < 42) || (-2.25 != float64(1.5))) {
                switch v3 {
                case "a":
                    v1 += 3
                case "foo bar":
                    v28 := []int{}
                    _ = v28
                }
                v29 := []int{fi27, fi27}
                var v30 int64 = -1
                _ = v29
                _ = v30
            }
            v25 = (v25 & (v25 >> ((255) & 63)))
            v25 += v25
            _ = fi27
        }
        v31 := GenS1{F0: v3}
        v32 := GenS4{F0: false, F1: -2147483648, F2: 0}
        _ = v31
        _ = v32
    }
    fmt.Println(v24, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
