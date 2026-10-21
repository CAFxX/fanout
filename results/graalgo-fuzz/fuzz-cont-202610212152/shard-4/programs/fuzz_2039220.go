package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 string }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int32 }
func genFn0(p0 int, p1 int16, p2 uint) int64 {
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint16 = 42
            _ = fi7
            _ = v8
        }
        v9 := []int{fi6, (p0 >> ((p2) & 63))}
        v9 = append(v9, -2147483648)
        _ = fi6
        _ = v9
    }
    var v10 uint64 = 0
    if (true && false) {
        for fi11 := 0; fi11 < 8; fi11++ {
            fmt.Println(p1, v10)
            switch (fi11 >> ((p2) & 63)) {
            case 4294967295:
                var v12 uint8 = 0
                _ = v12
            case 1:
                v13 := []int{(p0 - (fi11 >> (((p2 << ((p2) & 63))) & 63))), fi11, int(p1)}
                _ = v13
            case -2147483648:
                if true {
                    var v14 bool = false
                    var v15 string = ("\n\t" + "a")
                    _ = v14
                    _ = v15
                }
            }
            _ = fi11
        }
        var v16 uint8 = 26
        _ = v16
    } else {
        v17 := []int{p0}
        _ = v17
    }
    var v18 bool = ((!(!true)) && (!(!false)))
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v18
    return int64(p1)
}
func genFn1(p0 uint64) int64 {
    v19 := []int{2147483647, -1, -405, 2147483647}
    fmt.Println(p0)
    v20 := GenS1{F0: 0, F1: 9223372036854775807, F2: "foo bar"}
    _ = p0
    _ = v19
    _ = v20
    return genFn0(-228, -32768, 4294967295)
}

func main() {
    var v1 int64 = 564
    var v2 uint16 = 65535
    v3 := []string{"", (("" + ("\n\t" + "\n\t")) + "\n\t"), (("foo bar" + "") + "foo bar"), "a"}
    var v4 string = "a"
    v5 := make(map[uint32]int)
    if false {
        for fi21 := 0; fi21 < 5; fi21++ {
            if false {
                v4 = v4
                v22 := GenS2{F0: (true || true), F1: 128}
                var v23 int64 = -455
                _ = v22
                _ = v23
            } else {
                v24 := GenS3{F0: (v1 / ((v1) | 1)), F1: (v4 + v4)}
                _ = v24
            }
            _ = fi21
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v25 := GenS2{F0: (true && ((^(v2)) > v2)), F1: 0}
        _ = v25
    }
    var v26 float64 = 0.1
    for fi27 := 0; fi27 < 9; fi27++ {
        v28 := GenS4{F0: -1}
        v26 += v26
        v29 := GenS1{F0: 2, F1: (-(v1)), F2: "a"}
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := GenS1{F0: 42, F1: -1, F2: v4}
    var v31 int = 127
    v31 -= v31
    fmt.Println(v1)
    for fi32 := 0; fi32 < 12; fi32++ {
        var v33 int64 = (v1 >> ((34) & 63))
        v34 := v5[2147483648]
        var v35 int32 = -291
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    var v36 uint16 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v4
    _ = v5
}
