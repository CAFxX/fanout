package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int }
func genFn0(p0 uint32) uint {
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := GenS1{F0: 1}
        _ = fi4
        _ = v5
    }
    fmt.Println(p0)
    _ = p0
    return 508
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[string]bool)
    v3 := []string{("foo bar" + "a"), "hello"}
    switch -786 {
    case 128:
        v1 = (^(v1))
    case -9223372036854775808:
        for fi6 := 0; fi6 < 10; fi6++ {
            if (v1 < 2147483648) {
                v7 := GenS2{F0: -89}
                var v8 float32 = 1e10
                _ = v7
                _ = v8
            }
            fmt.Println(fi6, v1)
            v9 := GenS1{F0: -1}
            _ = fi6
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint = genFn0((^(v1)))
            if false {
                var v13 bool = (60 > 100)
                _ = v13
            }
            var v14 uint8 = 127
            if len(v3) > 0 {
                v3[0] = ""
            }
            v15 := GenS2{F0: 220}
            if (42 > 50) {
                v16 := GenS1{F0: 505}
                v1 = (v1 ^ (v1 << ((0) & 63)))
                _ = v16
            } else {
                if (v14 > v14) {
                    var v17 int64 = -529
                    _ = len(v3)
                    _ = v17
                } else {
                    v1 += v1
                    v18 := GenS3{F0: 2147483647}
                    _ = v18
                }
                v14 = v14
                v19 := GenS4{F0: 9223372036854775807}
                _ = v19
            }
            _ = fi11
            _ = v12
            _ = v14
            _ = v15
        }
        v20 := GenS2{F0: 42}
        _ = fi10
        _ = v20
    }
    v21 := GenS4{F0: 128}
    for fi22 := 0; fi22 < 9; fi22++ {
        v1 -= (^(v1))
        _ = fi22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        fmt.Println(v1, fi23)
        _ = fi23
    }
    v24 := GenS2{F0: 803}
    v3 = append(v3, ("foo bar" + "\n\t"))
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = 42
        switch -573 {
        case 32768:
            var v25 float64 = 0.1
            _ = v25
        case -9223372036854775808:
            v1 -= (v1 / ((v1) | 1))
        }
    }
    var v26 bool = false
    if v26 {
        for fi27 := 0; fi27 < 7; fi27++ {
            if v26 {
                v28 := GenS3{F0: -1}
                var v29 int16 = 1000
                if true {
                    v26 = ((v29 >> ((42) & 63)) >= 1)
                    v30 := v2["a"]
                    _ = v30
                }
                _ = v28
                _ = v29
            }
            _ = fi27
        }
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        v26 = (!(336 > -1))
        if (!v26) {
            for fi32 := 0; fi32 < 8; fi32++ {
                var v33 bool = (926 > 0)
                fmt.Println(v1)
                v33 = true
                if true {
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    var v34 int8 = 30
                    v34 += v34
                    v1 = v1
                    _ = v34
                }
                _ = fi32
                _ = v33
            }
        }
        _ = fi31
    }
    v1 = v1
    var v35 uint32 = (v1 ^ v1)
    var v36 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v26
    _ = v3
    _ = v35
    _ = v36
}
