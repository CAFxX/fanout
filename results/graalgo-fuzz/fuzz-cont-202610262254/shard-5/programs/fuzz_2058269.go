package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0() bool {
    var v6 int16 = -626
    v6 += v6
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 string = "hello"
        var v9 string = v8
        var v10 uint64 = 0
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := []int{1}
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 float32 = 1e10
        _ = fi12
        _ = v13
    }
    _ = v11
    _ = v6
    return (1.5 >= float32(1.5))
}
func genFn1(p0 uint8) uint {
    v14 := []int{1}
    switch 128 {
    case -292:
        var v15 int = -128
        _ = v15
    case -2147483648:
        if len(v14) > 0 {
            v14[0] = 32767
        }
    case -194:
        p0 = p0
    default:
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := GenS1{F0: 42}
            var v18 uint8 = p0
            fmt.Println(p0, fi16, p0)
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    v14 = append(v14, -1)
    p0 = (^(p0))
    p0 = 128
    _ = p0
    _ = v14
    return 710
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    var v3 uint32 = 2147483648
    v4 := []uint32{v3}
    var v5 uint8 = 127
    fmt.Println(v3, v5)
    var v19 int8 = 100
    switch 127 {
    case 128:
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 uint = 42
            v22 := GenS1{F0: 42}
            v23 := GenS1{F0: 642}
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
    default:
        for fi24 := 0; fi24 < 17; fi24++ {
            var v25 uint16 = 32767
            var v26 bool = true
            fmt.Println(v5, fi24)
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    if false {
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := GenS1{F0: 42}
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 3; fi29++ {
            v19 = v19
            fmt.Println(fi29, v19, v3)
            _ = fi29
        }
        var v30 uint64 = 18446744073709551615
        _ = v30
    }
    v3 -= v3
    var v31 int16 = -1
    if ((true && true) && ("" < "hello")) {
        var v32 float64 = -2.25
        var v33 bool = (!false)
        _ = v32
        _ = v33
    }
    var v34 int64 = 1
    v35 := v1["foo bar"]
    v31 += 1000
    v34 = (v34 * v34)
    v5 = 127
    var v36 int32 = -1
    switch -9223372036854775808 {
    case 32767:
        if false {
            v19 += 0
            if (false && true) {
                v37 := GenS2{F0: v35}
                var v38 int32 = (v36 / (((v36 / (((^(v36))) | 1))) | 1))
                _ = v37
                _ = v38
            } else {
                if false {
                    fmt.Println(v19, v31, v31)
                    v35 -= v35
                }
                v19 += v19
                v31 += v31
            }
        } else {
            v39 := GenS2{F0: 0}
            v19 = v19
            for fi40 := 0; fi40 < 20; fi40++ {
                if true {
                    switch (fi40 * (fi40 + 32768)) {
                    case -9223372036854775808:
                        for fi41 := 0; fi41 < 18; fi41++ {
                            v42 := GenS1{F0: 956}
                            var v43 int8 = -128
                            v3 -= v35
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                    default:
                        v44 := GenS1{F0: 0}
                        _ = v44
                    }
                    v4 = append(v4, (^(v35)))
                }
                _ = len(v4)
                _ = fi40
            }
            _ = v39
        }
    case 663:
        for fi45 := 0; fi45 < 17; fi45++ {
            v34 -= 9223372036854775807
            _ = fi45
        }
    }
    _ = len(v4)
    var v46 uint = 2
    v5 = (v5 << ((genFn1((^(v5)))) & 63))
    if len(v4) > 0 {
        v4[0] = 65535
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v46
    _ = v5
}
