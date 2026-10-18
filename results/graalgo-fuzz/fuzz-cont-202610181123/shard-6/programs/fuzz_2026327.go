package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int; F1 string; F2 uint64 }
func genFn0(p0 string, p1 bool, p2 uint32) int {
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int64 = 42
        var v9 int16 = int16(p2)
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    var v11 string = p0
    v12 := []int{301}
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 uint16 = 0
        for fi15 := 0; fi15 < 10; fi15++ {
            if p1 {
                v16 := GenS1{F0: 2147483647}
                _ = v16
            }
            v14 -= 608
            switch (fi13 * (fi15 + (fi15 - (fi15 % ((fi13) | 1))))) {
            case 0:
                fmt.Println(fi15, p1)
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -288
}

func main() {
    var v1 uint8 = 42
    var v2 uint16 = 32767
    var v3 int = -9223372036854775808
    var v4 float64 = 3.14159
    var v5 int16 = 1000
    var v6 bool = (v2 > (v2 | v2))
    var v17 uint16 = 854
    var v18 int16 = v5
    v5 += v18
    if (v5 > (v18 + (v18 & v18))) {
        var v19 uint = 866
        var v20 int16 = v5
        v21 := []int{(v3 + v3), -828}
        _ = v19
        _ = v20
        _ = v21
    } else {
        v6 = false
        var v22 uint8 = v1
        _ = v22
    }
    v23 := GenS1{F0: 0}
    if v6 {
        fmt.Println(v18)
    }
    if v6 {
        v24 := GenS1{F0: 62}
        _ = v24
    } else {
        if true {
            v25 := []int{v3, (^(v3)), 127, (v3 + (v3 / (((-(v3))) | 1)))}
            _ = v25
        }
        var v26 bool = v6
        v6 = false
        _ = v26
    }
    v27 := GenS1{F0: -1}
    v28 := GenS1{F0: -923}
    var v29 uint8 = v1
    if (((float32(float64(3.14159)) >= (0.1 * (3.14159 - float32(v4)))) && true) || (55 > v29)) {
        v30 := GenS1{F0: 369}
        switch 42 {
        case 32768:
            v31 := GenS2{F0: 1, F1: "\n\t"}
            _ = v31
        case 1:
            fmt.Println(v17)
        case 42:
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 uint = 2
                for fi34 := 0; fi34 < 11; fi34++ {
                    if v6 {
                        v33 -= 18446744073709551615
                        if false {
                            var v35 int64 = int64(fi32)
                            switch "" {
                            case "\n\t":
                                if v6 {
                                    v36 := []int{(fi32 & (v3 / ((int(v2)) | 1))), int(v35), genFn0("foo bar", (v6 && (4294967295 != uint64(v33))), 42)}
                                    if ((fi32 * (fi34 + fi32)) < genFn0("a", ((fi32 << ((v33) & 63)) != (fi34 >> ((v33) & 63))), 50)) {
                                        v18 = v5
                                        var v37 int = (-(fi32))
                                        var v38 uint32 = 471
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = v36
                                }
                            }
                            _ = v35
                        } else {
                            switch ("" + "foo bar") {
                            case "foo bar":
                                v39 := []int{int(v29), (v3 + 381), 127}
                                _ = v39
                            }
                        }
                        var v40 float32 = 1.5
                        v29 -= v1
                        if ((v33 - v33) < 418) {
                            switch "\n\t" {
                            case "\n\t":
                                v18 += 1
                            case "":
                                v4 -= v4
                            case "a":
                                fmt.Println(v6, fi32, v40)
                            default:
                                fmt.Println(v4, v29, fi32)
                            }
                            if (!((!false) && v6)) {
                                var v41 int16 = v18
                                v42 := GenS3{F0: 4294967295, F1: "\n\t", F2: 0}
                                v29 -= 2
                                _ = v41
                                _ = v42
                            }
                        }
                        _ = v40
                    }
                    _ = fi34
                }
                _ = fi32
                _ = v33
            }
        }
        _ = v30
    } else {
        if true {
            for fi43 := 0; fi43 < 8; fi43++ {
                var v44 string = ("" + "hello")
                v45 := []int{}
                var v46 uint = 65535
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
            }
        }
    }
    v3 -= (v3 + (v3 << ((18446744073709551615) & 63)))
    fmt.Println(v6, v3)
    v47 := GenS1{F0: -1}
    v48 := GenS2{F0: 42, F1: "hello"}
    var v49 int16 = (v18 - v18)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
