package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
func genFn0(p0 uint, p1 uint16) uint {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := []int{(fi4 << ((42) & 63))}
        fmt.Println(p0)
        if (!false) {
            var v6 uint64 = 1
            _ = v6
        } else {
            p1 = p1
            fmt.Println(p0)
        }
        _ = fi4
        _ = v5
    }
    fmt.Println(p1, p1)
    v7 := GenS1{F0: 42, F1: 1000}
    p0 -= (p0 - p0)
    _ = p0
    _ = p1
    _ = v7
    return 65535
}
func genFn1() uint8 {
    var v8 string = "a"
    switch (v8 + "\n\t") {
    case "a":
        for fi9 := 0; fi9 < 16; fi9++ {
            if ((1.5 < -2.25) || true) {
                var v10 bool = true
                v8 = v8
                _ = v10
            }
            switch (fi9 * (fi9 & -285)) {
            case 127:
                for fi11 := 0; fi11 < 7; fi11++ {
                    fmt.Println(v8, v8)
                    fmt.Println(v8, v8, fi11)
                    _ = fi11
                }
            }
            _ = fi9
        }
    case "foo bar":
        v8 = (("a" + v8) + "\n\t")
    case "\n\t":
        v12 := GenS1{F0: 4, F1: 0}
        _ = v12
    default:
        var v13 string = v8
        _ = v13
    }
    _ = v8
    return 127
}

func main() {
    var v1 bool = (970 == 65535)
    var v2 uint = 42
    v3 := []uint32{65535}
    v1 = v1
    var v14 uint8 = genFn1()
    v14 = v14
    fmt.Println(v1)
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int16 = 754
        for fi17 := 0; fi17 < 8; fi17++ {
            v1 = (false && v1)
            _ = len(v3)
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        v1 = ((^(v14)) == (^(v14)))
        var v19 int8 = 42
        _ = fi18
        _ = v19
    }
    var v20 bool = (!((v14 == 1) || (9223372036854775807 > 9223372036854775807)))
    switch 0 {
    case -28:
        if (524 >= 255) {
            for fi21 := 0; fi21 < 3; fi21++ {
                for fi22 := 0; fi22 < 18; fi22++ {
                    var v23 int32 = 2147483647
                    _ = fi22
                    _ = v23
                }
                v20 = false
                v14 += (v14 ^ 42)
                _ = fi21
            }
            fmt.Println(v2, v20)
        } else {
            for fi24 := 0; fi24 < 4; fi24++ {
                switch (fi24 & (fi24 ^ 1)) {
                case 127:
                    if v1 {
                        v1 = v1
                    } else {
                        var v25 uint32 = 908
                        v26 := GenS1{F0: 128, F1: -1}
                        _ = v25
                        _ = v26
                    }
                default:
                    fmt.Println(v14)
                }
                var v27 uint = (v2 | 65535)
                _ = fi24
                _ = v27
            }
            var v28 int16 = 1000
            var v29 float64 = 0.1
            _ = v28
            _ = v29
        }
    case -2147483648:
        v20 = ((!((("a" + "") + "") >= "hello")) || true)
    case -914:
        var v30 uint64 = 2
        _ = v30
    }
    var v31 int = -827
    switch "hello" {
    case "":
        var v32 string = (("a" + "foo bar") + "a")
        _ = v32
    case "foo bar":
        if v20 {
            for fi33 := 0; fi33 < 1; fi33++ {
                v34 := GenS1{F0: v31, F1: 32767}
                _ = fi33
                _ = v34
            }
            var v35 uint64 = 4294967295
            var v36 uint16 = 2
            _ = v35
            _ = v36
        }
    case "a":
        for fi37 := 0; fi37 < 19; fi37++ {
            if len(v3) > 0 {
                v3[0] = 42
            }
            if v20 {
                if (v20 && (!(!v1))) {
                    var v38 uint8 = 128
                    v14 = (v38 * v38)
                    if false {
                        var v39 int = (fi37 + 371)
                        var v40 uint64 = 914
                        v14 -= v38
                        _ = v39
                        _ = v40
                    }
                    _ = v38
                } else {
                    fmt.Println(v20, v1, fi37)
                }
                var v41 uint32 = 293
                var v42 uint32 = uint32(v2)
                _ = v41
                _ = v42
            }
            _ = fi37
        }
    }
    v14 = (v14 * (v14 + (v14 | genFn1())))
    var v43 bool = (v31 < v31)
    v1 = v20
    if (true || false) {
        var v44 float64 = 3.14159
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch (^(v31)) {
        case 128:
            var v45 int64 = 9223372036854775807
            _ = v45
        }
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v43
}
