package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint; F2 uint32 }
func genFn0(p0 int64) int64 {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 string = "foo bar"
        v8 := GenS1{F0: 41}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int64 = (-(p0))
    p0 = 1
    _ = p0
    _ = v9
    return p0
}
func genFn1(p0 uint16) int32 {
    var v10 int16 = -1
    if false {
        fmt.Println(v10, p0)
    } else {
        v11 := GenS1{F0: -128}
        _ = v11
    }
    _ = p0
    _ = v10
    return 996
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v5 uint16 = 32767
    v5 -= 2
    for fi12 := 0; fi12 < 4; fi12++ {
        if true {
            fmt.Println(fi12)
            delete(v2, "hello")
            if true {
                if false {
                    var v13 bool = true
                    var v14 float64 = 0.1
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 2; fi15++ {
                    switch fi15 {
                    case 626:
                        var v16 int = fi15
                        _ = v16
                    case -796:
                        v17 := GenS2{F0: 42, F1: 0, F2: 1}
                        _ = v17
                    default:
                        v18 := GenS2{F0: -548, F1: 42, F2: 65535}
                        _ = v18
                    }
                    _ = fi15
                }
                switch (("a" + "") + "") {
                case "foo bar":
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v20 := []int{(^(fi19))}
                        v21 := GenS1{F0: -128}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                case "":
                    v5 += (v5 & (v5 << ((65535) & 63)))
                }
            }
        }
        _ = fi12
    }
    fmt.Println(v5)
    switch -650 {
    case 42:
        v22 := GenS2{F0: 1, F1: 4294967295, F2: 2147483648}
        _ = v22
    case -447:
        if (0.0 != ((0.0 - 0.0) / 3.14159)) {
            if (false || (42 > 255)) {
                var v23 int64 = 1
                _ = v23
            }
        } else {
            var v24 int32 = genFn1(2)
            _ = v24
        }
    }
    if true {
        v2["hello"] = 42
        var v25 uint8 = 0
        _ = v25
    } else {
        fmt.Println(v5)
    }
    v26 := []int{}
    v5 -= v5
    v27 := GenS1{F0: 127}
    fmt.Println(v5)
    v5 += v5
    var v28 uint8 = 212
    if true {
        v29 := v4["foo bar"]
        if true {
            fmt.Println(v28)
            v29 += v29
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 uint8 = (^(v28))
                for fi32 := 0; fi32 < 16; fi32++ {
                    var v33 int16 = 686
                    _ = fi32
                    _ = v33
                }
                _ = fi30
                _ = v31
            }
        } else {
            var v34 uint64 = 163
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            _ = v34
        }
        fmt.Println(v5, v29)
        _ = v29
    } else {
        v35 := GenS2{F0: -333, F1: 12, F2: 4294967295}
        for fi36 := 0; fi36 < 13; fi36++ {
            var v37 uint64 = 2
            fmt.Println(v28, v5, v28)
            if true {
                delete(v4, "a")
            } else {
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
            }
            _ = fi36
            _ = v37
        }
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
