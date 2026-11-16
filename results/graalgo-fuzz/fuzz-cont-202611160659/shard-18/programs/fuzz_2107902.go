package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 bool; F1 uint; F2 uint }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int; F2 uint }
type GenS5 struct { F0 string }
func genFn0(p0 int8, p1 int8, p2 uint8) int16 {
    v5 := []int{524, 9223372036854775807, 127, 698}
    v5 = append(v5, -978)
    var v6 int8 = (-(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 string, p1 int32) int {
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            p0 = "foo bar"
            _ = fi8
        }
        p0 = p0
        _ = fi7
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 string = ((p0 + "\n\t") + "hello")
        p1 = p1
        p0 = ((p0 + "a") + v10)
        _ = fi9
        _ = v10
    }
    v11 := []int{-2147483648, 0, 127, -673}
    v11 = append(v11, -1)
    _ = p0
    _ = p1
    _ = v11
    return 32768
}

func main() {
    var v1 int64 = -256
    var v2 string = ("foo bar" + "hello")
    var v3 string = ""
    var v4 uint32 = 42
    fmt.Println(v1)
    v12 := GenS1{F0: v1, F1: v1}
    if false {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 int16 = 955
            _ = fi13
            _ = v14
        }
        if true {
            for fi15 := 0; fi15 < 20; fi15++ {
                for fi16 := 0; fi16 < 9; fi16++ {
                    if true {
                        var v17 int32 = 2147483647
                        v4 += (v4 >> ((4294967295) & 63))
                        v18 := []int{(fi16 & 32767)}
                        _ = v17
                        _ = v18
                    } else {
                        v19 := []int{-2147483648}
                        fmt.Println(fi16, v4, v3)
                        if ((!true) || (true || false)) {
                            v20 := GenS1{F0: (-(v1)), F1: -9223372036854775808}
                            v21 := GenS1{F0: v1, F1: (v1 + v1)}
                            _ = v20
                            _ = v21
                        }
                        _ = v19
                    }
                    var v22 uint8 = 255
                    _ = fi16
                    _ = v22
                }
                fmt.Println(v4, fi15, v2)
                _ = fi15
            }
        } else {
            if true {
                v23 := GenS2{F0: false, F1: 42, F2: 206}
                _ = v23
            } else {
                v1 -= v1
            }
        }
        v4 += (v4 >> ((4294967295) & 63))
    } else {
        for fi24 := 0; fi24 < 18; fi24++ {
            v3 = v3
            _ = fi24
        }
        fmt.Println(v1)
        fmt.Println(v2, v1)
    }
    v25 := GenS3{F0: 4294967295, F1: 0}
    v3 = ((v3 + (v3 + v2)) + "a")
    var v26 int8 = 42
    fmt.Println(v3)
    var v27 int64 = v1
    v4 = (^(v4))
    var v28 int16 = -32768
    var v29 uint64 = 583
    if true {
        v30 := []int{-2147483648, 4294967295}
        v4 = (v4 / (((v4 | (v4 << ((0) & 63)))) | 1))
        v30 = append(v30, 127)
        _ = v30
    }
    v31 := GenS4{F0: 65535, F1: 32767, F2: 255}
    var v32 float32 = 1e10
    v33 := GenS5{F0: (v3 + "hello")}
    var v34 int = 32768
    v4 = (v4 & v4)
    var v35 uint16 = 521
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
