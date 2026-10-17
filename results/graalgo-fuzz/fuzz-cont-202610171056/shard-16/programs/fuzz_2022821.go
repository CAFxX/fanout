package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint32) int8 {
    p0 += p0
    if true {
        fmt.Println(p1, p1)
        for fi6 := 0; fi6 < 18; fi6++ {
            if (true || false) {
                v7 := []int{(fi6 - fi6)}
                p1 = 0
                var v8 uint8 = 238
                _ = v7
                _ = v8
            } else {
                switch fi6 {
                case 897:
                    if ((!(!false)) && false) {
                        if (!false) {
                            fmt.Println(fi6)
                            var v9 int64 = (p0 * 42)
                            v10 := []int{}
                            _ = v10
                            _ = v9
                        }
                    } else {
                        for fi11 := 0; fi11 < 5; fi11++ {
                            if ((1 < 2) && true) {
                                var v12 uint16 = 981
                                var v13 bool = (!false)
                                v14 := GenS1{F0: 4294967295, F1: 4294967295, F2: 255}
                                v15 := GenS1{F0: 4294967295, F1: 2, F2: v12}
                                v16 := GenS2{F0: -485, F1: int8(p0)}
                                _ = v12
                                _ = v13
                                _ = v14
                                _ = v15
                                _ = v16
                            } else {
                                p0 += p0
                                switch (fi11 + (fi6 - int(p0))) {
                                case 422:
                                    var v17 int8 = 42
                                    _ = v17
                                default:
                                    p1 -= p1
                                }
                                v18 := GenS1{F0: 327, F1: p1, F2: 677}
                                _ = v18
                            }
                            fmt.Println(p1, fi6)
                            if true {
                                if (!(2 <= 4)) {
                                    p0 = (p0 - (p0 & p0))
                                }
                                v19 := []int{128}
                                var v20 uint64 = 0
                                switch ("" + "hello") {
                                case "":
                                    p1 += (^(p1))
                                }
                                switch fi6 {
                                case 42:
                                    v21 := GenS1{F0: 344, F1: 0, F2: 32767}
                                    _ = v21
                                case 1:
                                    v20 += uint64(p1)
                                }
                                _ = v19
                                _ = v20
                            }
                            _ = fi11
                        }
                    }
                case -2147483648:
                    fmt.Println(p0)
                case 699:
                    v22 := []int{fi6, -627}
                    _ = v22
                }
                p1 = 2147483647
            }
            _ = fi6
        }
        v23 := []int{-1, -1}
        _ = v23
    } else {
        switch ("\n\t" + "") {
        case "a":
            v24 := GenS2{F0: -1, F1: -128}
            _ = v24
        }
        v25 := GenS2{F0: 0, F1: 1}
        _ = v25
    }
    fmt.Println(p0)
    fmt.Println(p0)
    var v26 uint16 = 65535
    _ = p0
    _ = p1
    _ = v26
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 429
    v3 := []uint32{}
    var v4 int16 = -1
    v5 := make(map[int]int)
    v27 := GenS2{F0: v2, F1: 0}
    if (false && (65535 == 797)) {
        for fi28 := 0; fi28 < 3; fi28++ {
            for fi29 := 0; fi29 < 10; fi29++ {
                delete(v1, 2147483647)
                v30 := v5[(-(fi29))]
                _ = fi29
                _ = v30
            }
            var v31 uint8 = 0
            var v32 float32 = 1e10
            _ = fi28
            _ = v31
            _ = v32
        }
        fmt.Println(v2, v2)
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        if true {
            var v34 bool = true
            _ = v34
        } else {
            v35 := GenS2{F0: (v2 | v2), F1: -128}
            _ = v35
        }
        var v36 int8 = 1
        _ = fi33
        _ = v36
    }
    var v37 uint16 = 0
    v37 -= v37
    for fi38 := 0; fi38 < 3; fi38++ {
        switch (^(fi38)) {
        case 32768:
            v37 = v37
        }
        v39 := GenS3{F0: v37}
        v40 := v5[fi38]
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 int16 = (v4 << ((4294967295) & 63))
    fmt.Println(v4, v41, v2)
    v37 = (v37 | (v37 ^ v37))
    v42 := GenS3{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
