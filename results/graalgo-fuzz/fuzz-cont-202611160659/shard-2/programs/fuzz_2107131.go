package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 string }
type GenS2 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS5 struct { F0 string }

func main() {
    var v1 uint8 = 1
    var v2 float32 = 0.0
    v3 := make(map[uint32]bool)
    var v4 int64 = 615
    var v5 int32 = -951
    v1 += v1
    if false {
        var v6 int = 4294967295
        var v7 float32 = (v2 - (float32(((float64(v2) - -2.25) - 1.5)) + v2))
        if false {
            v6 = v6
            var v8 int16 = int16(v1)
            v9 := []int{-1, (v6 & (v6 & v6)), -576, v6}
            _ = v8
            _ = v9
        } else {
            var v10 uint = 42
            _ = v10
        }
        _ = v6
        _ = v7
    } else {
        if (!true) {
            for fi11 := 0; fi11 < 11; fi11++ {
                fmt.Println(v2, v4)
                if true {
                    v12 := v3[42]
                    _ = v12
                }
                _ = fi11
            }
            v5 = 0
            var v13 int64 = (v4 - v4)
            _ = v13
        } else {
            switch ("hello" + "") {
            case "":
                for fi14 := 0; fi14 < 13; fi14++ {
                    if ((18446744073709551615 > 18446744073709551615) || false) {
                        v15 := GenS1{F0: fi14, F1: v4, F2: "foo bar"}
                        fmt.Println(v2)
                        v16 := GenS2{F0: 799, F1: 4294967295, F2: v4}
                        _ = v15
                        _ = v16
                    } else {
                        var v17 int64 = (v4 >> ((81) & 63))
                        v18 := GenS2{F0: 18446744073709551615, F1: 1, F2: v17}
                        v4 += (v4 ^ v17)
                        _ = v17
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 2; fi19++ {
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v1)
                        }
                        v20 := GenS1{F0: (fi19 ^ 974), F1: (v4 & v4), F2: "hello"}
                        _ = fi19
                        _ = v20
                    }
                    _ = fi14
                }
            case "foo bar":
                v5 -= -954
            case "a":
                var v21 bool = (true && (1 != 32767))
                _ = v21
            }
            var v22 int16 = -32768
            var v23 int16 = v22
            _ = v22
            _ = v23
        }
        fmt.Println(v2)
        fmt.Println(v4, v1)
    }
    v24 := GenS3{F0: 42, F1: "foo bar"}
    var v25 uint16 = 255
    if false {
        v2 = (0.1 + v2)
        v5 -= v5
    }
    v5 = (v5 >> ((65535) & 63))
    v26 := GenS3{F0: v25, F1: ""}
    v27 := GenS4{F0: 790, F1: 966, F2: -1}
    v4 = v4
    if (((^(v5)) == 1) && (!(!false))) {
        var v28 string = "\n\t"
        v29 := v3[2]
        _ = v28
        _ = v29
    }
    if (false && (true && false)) {
        v4 += int64(v5)
        v25 = (v25 / ((v25) | 1))
    }
    v4 = v4
    var v30 bool = false
    v31 := GenS2{F0: 42, F1: 4294967295, F2: -9223372036854775808}
    var v32 bool = v30
    v33 := GenS5{F0: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
