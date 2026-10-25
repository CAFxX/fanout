package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint16 }
func genFn0(p0 uint8, p1 bool, p2 bool) bool {
    switch "\n\t" {
    case "a":
        v6 := GenS1{F0: "", F1: 0, F2: 0}
        _ = v6
    case "":
        var v7 int = 32767
        _ = v7
    default:
        v8 := GenS2{F0: 4294967295}
        _ = v8
    }
    v9 := []int{}
    for fi10 := 0; fi10 < 1; fi10++ {
        if p1 {
            var v11 uint8 = (p0 & p0)
            _ = v11
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (32767 == 0)
}

func main() {
    var v1 bool = (false || (127 != 128))
    var v2 float32 = -2.25
    var v3 int8 = 1
    var v4 uint32 = 42
    v5 := []int{2147483647, -1}
    _ = len(v5)
    v1 = v1
    fmt.Println(v2)
    v1 = (!v1)
    if (!((v4 << ((185) & 63)) >= v4)) {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 int8 = (-(v3))
            if v1 {
                if (595 > 42) {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v2 -= v2
                        _ = fi14
                    }
                    if (fi12 == (-(fi12))) {
                        if (!false) {
                            fmt.Println(v2, v1)
                            v1 = (false && (805 <= 42))
                            var v15 int64 = 0
                            _ = v15
                        }
                    } else {
                        for fi16 := 0; fi16 < 12; fi16++ {
                            _ = len(v5)
                            v5 = append(v5, fi12)
                            fmt.Println(v3, v13)
                            _ = fi16
                        }
                    }
                } else {
                    v5 = append(v5, fi12)
                }
            } else {
                var v17 uint = 804
                if v1 {
                    var v18 string = ""
                    v5 = append(v5, (fi12 * (fi12 + (fi12 >> ((157) & 63)))))
                    _ = v18
                }
                _ = v17
            }
            _ = fi12
            _ = v13
        }
        v1 = (!((float64((0.1 - (v2 + v2))) + 0.0) <= 0.0))
    } else {
        var v19 int16 = 0
        v20 := GenS3{F0: 255}
        _ = v19
        _ = v20
    }
    v4 = v4
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v21 := GenS4{F0: true, F1: 42, F2: 154}
    fmt.Println(v1, v2, v1)
    v2 = 3.14159
    if true {
        fmt.Println(v2, v3)
        v22 := GenS2{F0: 9223372036854775807}
        for fi23 := 0; fi23 < 20; fi23++ {
            fmt.Println(v2, fi23)
            v3 = v3
            v24 := GenS2{F0: (fi23 << ((844) & 63))}
            _ = fi23
            _ = v24
        }
        _ = v22
    } else {
        if v1 {
            v5 = append(v5, -2147483648)
        }
        v3 = v3
    }
    v25 := GenS3{F0: 32767}
    var v26 uint = 0
    v3 -= v3
    fmt.Println(v1)
    v1 = v1
    v27 := GenS3{F0: 255}
    var v28 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
