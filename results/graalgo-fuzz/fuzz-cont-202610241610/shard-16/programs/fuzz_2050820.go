package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
func genFn0(p0 string) int64 {
    p0 = (p0 + "a")
    if (true || (!(66 == 32767))) {
        if true {
            v7 := []int{-796, 2147483647, -2147483648}
            for fi8 := 0; fi8 < 4; fi8++ {
                if true {
                    if false {
                        p0 = "a"
                        if (!(-781 == -9223372036854775808)) {
                            p0 = p0
                        } else {
                            for fi9 := 0; fi9 < 18; fi9++ {
                                p0 = (p0 + p0)
                                fmt.Println(fi9)
                                p0 = ((p0 + p0) + "\n\t")
                                _ = fi9
                            }
                        }
                    }
                    v10 := GenS1{F0: 0, F1: 2147483647}
                    if (((fi8 / ((fi8) | 1)) == (fi8 >> ((1) & 63))) && (("a" >= (p0 + (p0 + ((p0 + p0) + (p0 + (p0 + "foo bar")))))) || true)) {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        if false {
                            var v11 int16 = 42
                            var v12 uint = 720
                            _ = v11
                            _ = v12
                        } else {
                            for fi13 := 0; fi13 < 9; fi13++ {
                                v14 := GenS1{F0: 340, F1: 42}
                                _ = fi13
                                _ = v14
                            }
                            v15 := GenS1{F0: 1, F1: 4294967295}
                            var v16 int8 = 127
                            _ = v15
                            _ = v16
                        }
                    }
                    _ = v10
                } else {
                    var v17 int8 = 31
                    var v18 uint64 = 157
                    p0 = (p0 + p0)
                    _ = v17
                    _ = v18
                }
                _ = fi8
            }
            _ = v7
        }
    } else {
        fmt.Println(p0)
        var v19 uint8 = 181
        if (!(!true)) {
            for fi20 := 0; fi20 < 3; fi20++ {
                fmt.Println(p0, fi20)
                p0 = (p0 + ((p0 + p0) + (p0 + p0)))
                var v21 bool = ((-2.25 + ((1e10 - (0.1 - 3.14159)) * (-2.25 - float64((-2.25 + 0.1))))) == float64(0.0))
                _ = fi20
                _ = v21
            }
        } else {
            p0 = "foo bar"
            v22 := []int{1, 9223372036854775807}
            var v23 uint16 = 0
            _ = v22
            _ = v23
        }
        _ = v19
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        p0 = p0
        _ = fi24
    }
    _ = p0
    return -1
}
func genFn1() int16 {
    var v25 int16 = -846
    var v26 int8 = -31
    var v27 float32 = -2.25
    v28 := []int{-2147483648, -2147483648, 1}
    v29 := GenS1{F0: -874, F1: 2147483648}
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    return (v25 >> ((65535) & 63))
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]string)
    v3 := []uint32{2147483648, 122, 2, 0}
    var v4 int = 32768
    v5 := []int{(v4 % (((^(v4))) | 1)), (v4 & v4), (v4 << ((843) & 63)), 32768}
    var v6 string = "\n\t"
    var v30 uint8 = 42
    v31 := GenS1{F0: 2147483647, F1: 0}
    for fi32 := 0; fi32 < 20; fi32++ {
        var v33 bool = (145 != 2)
        v1 = genFn1()
        var v34 string = v6
        _ = fi32
        _ = v33
        _ = v34
    }
    v4 = (v4 ^ (v4 + 32767))
    _ = len(v2)
    fmt.Println(v1, v6)
    var v35 int16 = (v1 << ((18446744073709551615) & 63))
    delete(v2, 65535)
    v30 = 127
    v1 += (v35 | v1)
    v36 := GenS1{F0: 2147483647, F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
