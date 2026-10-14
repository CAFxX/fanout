package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 string }
type GenS2 struct { F0 int }
func genFn0(p0 uint32, p1 int16) int16 {
    var v4 bool = true
    if (65535 > 222) {
        v5 := GenS1{F0: (-(p1)), F1: 1, F2: ""}
        v4 = (!(false && (-1 >= 0)))
        _ = v5
    }
    v6 := []int{-678, 0}
    p1 += p1
    if v4 {
        _ = len(v6)
    } else {
        var v7 int64 = -9223372036854775808
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return (p1 | (p1 >> ((48) & 63)))
}
func genFn1(p0 uint64, p1 int, p2 uint64) uint16 {
    p0 += p2
    p2 = 266
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 int = 128
    var v2 string = ("" + "\n\t")
    v3 := []string{}
    v8 := GenS1{F0: 0, F1: 146, F2: v2}
    if len(v3) > 0 {
        v3[0] = v2
    }
    v1 = (-(v1))
    var v9 uint16 = 42
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: int16(v9), F1: 423, F2: v2}
        _ = len(v3)
        _ = fi10
        _ = v11
    }
    if (true || (!(false && (1 <= 536)))) {
        switch (^(v1)) {
        case 4294967295:
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        }
        var v12 int16 = genFn0(0, -1)
        _ = v12
    } else {
        if true {
            switch (v2 + "") {
            case "\n\t":
                v9 = (v9 - 255)
            case "a":
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 int = (v1 << ((596) & 63))
                    fmt.Println(fi13, v1, v2)
                    _ = fi13
                    _ = v14
                }
            case "hello":
                _ = len(v3)
            }
            _ = len(v3)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            var v15 float32 = ((-2.25 - 0.0) + float32(((-2.25 + 0.1) - 3.14159)))
            _ = v15
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            for fi17 := 0; fi17 < 19; fi17++ {
                v9 -= 2
                var v18 uint32 = 351
                _ = fi17
                _ = v18
            }
            v1 -= v1
            _ = fi16
        }
        switch (v2 + "") {
        case "foo bar":
            for fi19 := 0; fi19 < 8; fi19++ {
                for fi20 := 0; fi20 < 20; fi20++ {
                    v9 += (v9 << ((18446744073709551615) & 63))
                    fmt.Println(fi19, v9, v9)
                    fmt.Println(v9, fi19, v1)
                    _ = fi20
                }
                _ = len(v3)
                _ = fi19
            }
        case "a":
            var v21 bool = true
            _ = v21
        default:
            var v22 int32 = 0
            _ = v22
        }
    }
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 int16 = 1000
            v9 = v9
            if true {
                fmt.Println(fi23)
            }
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 2; fi25++ {
            if (338 == 0) {
                var v26 int16 = -1
                var v27 bool = (!(42 != 100))
                _ = v26
                _ = v27
            }
            _ = fi25
        }
    } else {
        v9 += (v9 * 255)
        if true {
            if false {
                var v28 int64 = -1
                for fi29 := 0; fi29 < 19; fi29++ {
                    var v30 uint32 = 0
                    _ = fi29
                    _ = v30
                }
                _ = v28
            }
            if (1 <= -145) {
                v2 = "hello"
            } else {
                v1 = v1
                switch v1 {
                case 4294967295:
                    var v31 float64 = (1e10 - (float64(0.0) + 1e10))
                    _ = v31
                }
            }
            v32 := GenS1{F0: 1, F1: 4294967295, F2: "foo bar"}
            _ = v32
        }
        switch v2 {
        case "":
            if (9223372036854775807 >= 42) {
                var v33 float32 = -2.25
                var v34 uint = 42
                var v35 uint64 = 42
                _ = v33
                _ = v34
                _ = v35
            }
        default:
            if len(v3) > 0 {
                v3[0] = v2
            }
        }
    }
    var v36 uint64 = 350
    v37 := GenS2{F0: (v1 % (((v1 ^ (v1 ^ v1))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v8
    _ = v9
}
