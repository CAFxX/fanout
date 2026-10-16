package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int64; F1 int16 }
func genFn0(p0 bool) int {
    switch "foo bar" {
    case "hello":
        switch "\n\t" {
        case "":
            for fi6 := 0; fi6 < 3; fi6++ {
                fmt.Println(fi6, p0)
                _ = fi6
            }
        case "a":
            if (!false) {
                for fi7 := 0; fi7 < 4; fi7++ {
                    fmt.Println(fi7, p0)
                    _ = fi7
                }
                var v8 int64 = 9223372036854775807
                fmt.Println(p0, p0)
                _ = v8
            } else {
                var v9 int32 = 2147483647
                v10 := GenS1{F0: (276 != 2), F1: 2147483647, F2: 1}
                if ((p0 && p0) || true) {
                    if p0 {
                        var v11 string = "hello"
                        _ = v11
                    }
                }
                _ = v10
                _ = v9
            }
        case "foo bar":
            if (1 > 627) {
                var v12 uint32 = 1
                v12 += 2147483648
                fmt.Println(v12, p0)
                _ = v12
            }
        }
    }
    v13 := GenS2{F0: 32767, F1: 716, F2: 0}
    _ = p0
    _ = v13
    return 4294967295
}

func main() {
    var v1 float64 = 1e10
    var v2 uint32 = 1
    var v3 uint = 739
    var v4 uint8 = 42
    var v5 float32 = 0.0
    v3 = 708
    var v14 int8 = -1
    var v15 int8 = (-(v14))
    v1 -= 3.14159
    fmt.Println(v4, v5, v14)
    var v16 bool = (-2147483648 >= -1)
    switch genFn0((324 == 4294967295)) {
    case 22:
        if (127 >= 127) {
            var v17 uint16 = uint16(v4)
            _ = v17
        } else {
            switch -9223372036854775808 {
            case 2147483647:
                var v18 int = 32768
                _ = v18
            }
            switch "a" {
            case "hello":
                v19 := GenS1{F0: (!v16), F1: v2, F2: (^(v14))}
                _ = v19
            case "foo bar":
                var v20 int8 = (v14 << ((v3) & 63))
                _ = v20
            }
        }
    }
    v21 := GenS3{F0: 1000, F1: "foo bar"}
    switch 2147483647 {
    case 0:
        var v22 int = 32767
        _ = v22
    case 32767:
        var v23 int16 = -1
        _ = v23
    case 128:
        if v16 {
            for fi24 := 0; fi24 < 1; fi24++ {
                for fi25 := 0; fi25 < 17; fi25++ {
                    switch "\n\t" {
                    case "hello":
                        v26 := GenS4{F0: 1, F1: 32767}
                        v27 := []int{-853}
                        if v16 {
                            var v28 string = "foo bar"
                            var v29 int64 = int64(v2)
                            v14 = 0
                            _ = v28
                            _ = v29
                        } else {
                            v14 += v14
                            var v30 int8 = (v15 | (v15 - v15))
                            _ = v30
                        }
                        var v31 uint8 = (v4 + (v4 | v4))
                        _ = v26
                        _ = v27
                        _ = v31
                    }
                    var v32 float64 = v1
                    _ = fi25
                    _ = v32
                }
                var v33 uint32 = v2
                _ = fi24
                _ = v33
            }
            v14 -= (v14 & v15)
        } else {
            var v34 uint32 = v2
            v35 := []int{32768, -1, 53, 4294967295}
            _ = v34
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
