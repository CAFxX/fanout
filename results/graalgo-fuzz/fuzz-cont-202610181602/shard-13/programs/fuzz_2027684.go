package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 string, p1 int32, p2 int32) uint16 {
    p2 = 42
    if true {
        p1 += 2147483647
        v4 := GenS1{F0: 42}
        _ = v4
    }
    fmt.Println(p2, p2, p1)
    p0 = (p0 + "")
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint = 135
    var v2 bool = false
    v3 := make(map[uint32]int)
    if v2 {
        v2 = true
    } else {
        v5 := []int{-659, 32767, -645, 42}
        _ = v5
    }
    if false {
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := GenS1{F0: 1}
            if ((!(v2 || (42 == -385))) || (!v2)) {
                var v8 uint8 = 255
                _ = v8
            } else {
                v9 := []int{(fi6 % ((fi6) | 1))}
                fmt.Println(fi6, fi6, v1)
                _ = v9
            }
            fmt.Println(v2, fi6, fi6)
            _ = fi6
            _ = v7
        }
    } else {
        if v2 {
            var v10 string = "\n\t"
            for fi11 := 0; fi11 < 14; fi11++ {
                v2 = false
                var v12 uint32 = 0
                _ = fi11
                _ = v12
            }
            _ = v10
        } else {
            v13 := []int{742}
            _ = v13
        }
        var v14 int64 = 1
        _ = v14
    }
    _ = len(v3)
    if v2 {
        fmt.Println(v2, v1)
        v2 = v2
        v15 := v3[856]
        _ = v15
    }
    var v16 int32 = 0
    fmt.Println(v16, v2)
    fmt.Println(v2, v2)
    v2 = v2
    switch 0 {
    case 0:
        if (-2147483648 != 1) {
            v2 = (127 >= 909)
            var v17 uint8 = 42
            _ = v17
        }
    }
    v2 = v2
    var v18 uint8 = 127
    v19 := []int{432, 0}
    if true {
        for fi20 := 0; fi20 < 4; fi20++ {
            if (-1 == 127) {
                var v21 int64 = 0
                if v2 {
                    v22 := GenS1{F0: -892}
                    v1 += (v1 / (((v1 - v1)) | 1))
                    switch fi20 {
                    case -496:
                        for fi23 := 0; fi23 < 6; fi23++ {
                            v16 -= (v16 & v16)
                            fmt.Println(v21, fi23)
                            var v24 int8 = 1
                            if false {
                                v21 = v21
                                v16 += (v16 | (v16 & (v16 + (v16 ^ v16))))
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                                v25 := GenS1{F0: -32768}
                                _ = v25
                            }
                            fmt.Println(fi20, fi20, fi23)
                            _ = fi23
                            _ = v24
                        }
                    case 4294967295:
                        v1 = v1
                    case 0:
                        var v26 bool = v2
                        _ = v26
                    default:
                        v2 = v2
                    }
                    _ = v22
                } else {
                    delete(v3, 582)
                }
                v18 += 2
                _ = v21
            }
            _ = fi20
        }
    } else {
        switch 127 {
        case 32768:
            fmt.Println(v16, v2, v1)
        }
    }
    var v27 uint64 = 0
    var v28 int16 = 1
    _ = len(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
