package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 int; F2 bool }
func genFn0(p0 int16, p1 uint) string {
    v5 := []int{32767, 0, -9223372036854775808, int(p1)}
    var v6 uint32 = 65535
    var v7 uint = 72
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return ""
}
func genFn1(p0 int, p1 uint, p2 uint64) int64 {
    if true {
        switch p0 {
        case 32768:
            v8 := GenS1{F0: true}
            _ = v8
        case -980:
            var v9 int64 = -9223372036854775808
            _ = v9
        case -1:
            p1 += 723
        }
    }
    v10 := GenS1{F0: true}
    v11 := GenS1{F0: (262 < 42)}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 157
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]int)
    var v3 string = (v1 + "foo bar")
    var v4 bool = (!(true || (!(!(!false)))))
    fmt.Println(v1, v1, v4)
    var v12 string = "a"
    var v13 uint16 = 255
    if (1 >= 337) {
        for fi14 := 0; fi14 < 1; fi14++ {
            v1 = v12
            _ = len(v2)
            _ = fi14
        }
        switch ("a" + v3) {
        case "foo bar":
            if (0 > 599) {
                v2[1] = 32768
                var v15 int8 = -80
                _ = v15
            } else {
                v2[0] = 2147483647
                var v16 int64 = genFn1(-1, 190, 4294967295)
                _ = v16
            }
        }
    } else {
        v4 = v4
        for fi17 := 0; fi17 < 1; fi17++ {
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 int64 = -9223372036854775808
                var v20 string = v3
                _ = fi18
                _ = v19
                _ = v20
            }
            v21 := []int{(fi17 | 9223372036854775807), fi17, (^(fi17)), (fi17 % ((fi17) | 1))}
            _ = fi17
            _ = v21
        }
    }
    if ((!((9223372036854775807 <= 0) || v4)) && true) {
        if (1.5 == 1e10) {
            v22 := GenS1{F0: v4}
            switch -268 {
            case 482:
                for fi23 := 0; fi23 < 12; fi23++ {
                    v24 := GenS1{F0: v4}
                    v25 := []int{}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            case 330:
                for fi26 := 0; fi26 < 19; fi26++ {
                    if (42 >= -743) {
                        v27 := GenS1{F0: v4}
                        switch v1 {
                        case "\n\t":
                            for fi28 := 0; fi28 < 5; fi28++ {
                                v29 := GenS2{F0: 461, F1: v13}
                                v1 = v3
                                _ = fi28
                                _ = v29
                            }
                        case "foo bar":
                            var v30 int16 = 185
                            _ = v30
                        }
                        _ = v27
                    }
                    _ = fi26
                }
            case -9223372036854775808:
                var v31 float64 = 1e10
                _ = v31
            }
            var v32 uint = 42
            _ = v22
            _ = v32
        }
    }
    var v33 uint64 = 4294967295
    if (-391 == 42) {
        fmt.Println(v1)
        v34 := []int{int(v13), 361}
        v35 := GenS3{F0: 0}
        _ = v34
        _ = v35
    } else {
        var v36 int16 = -32768
        if v4 {
            v37 := GenS3{F0: -92}
            _ = v37
        }
        var v38 uint8 = 4
        _ = v36
        _ = v38
    }
    v13 = v13
    v39 := GenS4{F0: (v12 + "a"), F1: 42, F2: (255 == 1)}
    if v4 {
        for fi40 := 0; fi40 < 14; fi40++ {
            switch v12 {
            case "a":
                if (v13 != v13) {
                    var v41 string = (("hello" + v12) + (v3 + (v3 + ("" + v12))))
                    var v42 int64 = 0
                    for fi43 := 0; fi43 < 19; fi43++ {
                        delete(v2, 42)
                        _ = fi43
                    }
                    _ = v41
                    _ = v42
                }
            }
            _ = fi40
        }
        fmt.Println(v33)
    } else {
        v44 := GenS1{F0: false}
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v39
    _ = v4
}
