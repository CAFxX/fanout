package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int64; F2 int8 }
func genFn0() uint {
    switch (("hello" + "") + ("\n\t" + "\n\t")) {
    case "foo bar":
        var v4 uint = 2
        _ = v4
    case "hello":
        v5 := []int{-9223372036854775808}
        _ = v5
    }
    var v6 bool = (false || true)
    _ = v6
    return 255
}
func genFn1(p0 string, p1 int64) uint16 {
    fmt.Println(p1, p1)
    var v7 int32 = 1
    _ = p0
    _ = p1
    _ = v7
    return 2
}

func main() {
    var v1 string = "a"
    var v2 uint64 = 870
    v3 := make(map[string]int)
    if false {
        v1 = "\n\t"
    }
    v1 = "a"
    if true {
        var v8 uint8 = 157
        _ = v8
    } else {
        for fi9 := 0; fi9 < 2; fi9++ {
            for fi10 := 0; fi10 < 18; fi10++ {
                if (1 == 1) {
                    if true {
                        if false {
                            v11 := GenS1{F0: 1}
                            v2 = (v2 - v2)
                            _ = v11
                        }
                    }
                    if false {
                        v1 = "\n\t"
                        var v12 int8 = -24
                        v13 := GenS1{F0: 1}
                        _ = v12
                        _ = v13
                    }
                    switch fi10 {
                    case 2147483647:
                        var v14 uint8 = 255
                        delete(v3, v1)
                        _ = v14
                    case -599:
                        var v15 uint32 = 0
                        var v16 float64 = float64(-2.25)
                        _ = v15
                        _ = v16
                    case 32767:
                        fmt.Println(fi10, v1, fi9)
                    }
                } else {
                    v1 = v1
                    switch (fi9 - (fi9 / (((fi9 | fi9)) | 1))) {
                    case 358:
                        fmt.Println(v2)
                    case 4294967295:
                        var v17 bool = (!((42 < 42) || (-1 == 42)))
                        _ = v17
                    default:
                        var v18 int = fi10
                        _ = v18
                    }
                }
                _ = len(v3)
                var v19 bool = (!(("a" != v1) || (!false)))
                _ = fi10
                _ = v19
            }
            var v20 uint64 = (v2 >> ((857) & 63))
            if false {
                fmt.Println(fi9, v2, v20)
                v1 = v1
                v3["foo bar"] = (fi9 << ((genFn0()) & 63))
            } else {
                v1 = v1
            }
            _ = fi9
            _ = v20
        }
    }
    v1 = (v1 + "")
    fmt.Println(v1, v1)
    for fi21 := 0; fi21 < 10; fi21++ {
        v1 = (v1 + v1)
        _ = fi21
    }
    var v22 uint64 = (v2 % ((v2) | 1))
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 uint32 = 42
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: 0}
    v2 = (v22 - 938)
    var v26 uint32 = 2
    v27 := []int{}
    if (!(9223372036854775807 >= 42)) {
        v22 += (v2 & v2)
        v28 := GenS1{F0: 568}
        _ = v28
    } else {
        if true {
            _ = len(v27)
            var v29 string = v1
            switch "foo bar" {
            case "":
                v22 -= v22
            case "hello":
                v30 := GenS1{F0: -2147483648}
                _ = v30
            default:
                for fi31 := 0; fi31 < 17; fi31++ {
                    var v32 uint32 = (v26 << ((18446744073709551615) & 63))
                    _ = fi31
                    _ = v32
                }
            }
            _ = v29
        } else {
            v26 = uint32(v22)
        }
        if true {
            var v33 int64 = 0
            for fi34 := 0; fi34 < 20; fi34++ {
                v35 := GenS2{F0: fi34, F1: (v33 | (v33 >> ((647) & 63))), F2: 127}
                _ = fi34
                _ = v35
            }
            _ = v33
        }
    }
    if false {
        var v36 int32 = -129
        _ = v36
    }
    var v37 bool = false
    v38 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v38
}
