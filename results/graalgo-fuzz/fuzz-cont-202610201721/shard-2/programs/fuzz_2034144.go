package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint32 }
func genFn0(p0 uint64, p1 string) uint8 {
    v6 := []int{32768, 32767}
    var v7 int64 = 9223372036854775807
    v7 += v7
    p0 = 560
    var v8 uint8 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int16 = 1
    v2 := make(map[string]int)
    var v3 int = 128
    v4 := make(map[int]bool)
    var v5 uint32 = 65535
    fmt.Println(v3)
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint16 = 2
        _ = fi9
        _ = v10
    }
    if (1 > 0) {
        var v11 string = "\n\t"
        v11 = v11
        _ = v11
    } else {
        switch -9223372036854775808 {
        case 568:
            v1 += -32768
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS1{F0: -492}
            v1 = (v1 ^ (v1 + (v1 + v1)))
            for fi14 := 0; fi14 < 19; fi14++ {
                if true {
                    delete(v2, "\n\t")
                } else {
                    v5 = v5
                }
                v15 := GenS2{F0: genFn0(42, ("hello" + "\n\t")), F1: uint64(v5), F2: 767}
                v16 := []int{4294967295, (-(fi12)), 42, 128}
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
        if true {
            for fi17 := 0; fi17 < 13; fi17++ {
                switch (fi17 + (v3 - (fi17 * v3))) {
                case -583:
                    for fi18 := 0; fi18 < 4; fi18++ {
                        if (127 <= 1) {
                            v19 := GenS3{F0: "", F1: -1}
                            v20 := GenS4{F0: 975}
                            if false {
                                v1 -= (^(v1))
                                delete(v4, -504)
                                v21 := []int{fi18, (fi18 * int(v5)), -384}
                                _ = v21
                            }
                            _ = v19
                            _ = v20
                        }
                        _ = fi18
                    }
                case -249:
                    var v22 uint = 10
                    _ = v22
                case 88:
                    var v23 uint64 = 70
                    _ = v23
                }
                _ = fi17
            }
        } else {
            v24 := GenS5{F0: 184, F1: 1, F2: v5}
            _ = v24
        }
    }
    var v25 uint32 = v5
    v26 := []int{}
    if ((false || (17 == 168)) && false) {
        switch v3 {
        case 0:
            var v27 uint32 = 4294967295
            _ = v27
        }
    }
    if ("" > "hello") {
        v1 = v1
    }
    v28 := GenS4{F0: 127}
    switch v3 {
    case 42:
        v29 := GenS1{F0: 42}
        _ = v29
    case 128:
        switch "foo bar" {
        case "hello":
            v3 = v3
        case "a":
            if (v1 < 762) {
                for fi30 := 0; fi30 < 10; fi30++ {
                    v31 := GenS1{F0: 0}
                    if false {
                        var v32 int16 = 1000
                        _ = v32
                    }
                    _ = fi30
                    _ = v31
                }
                var v33 uint = 18446744073709551615
                _ = v33
            } else {
                v25 = v25
                fmt.Println(v3)
                v25 += (v5 << ((618) & 63))
            }
        }
    case 32767:
        v1 -= v1
    default:
        v1 -= v1
    }
    var v34 int16 = -1
    for fi35 := 0; fi35 < 17; fi35++ {
        v36 := GenS4{F0: v3}
        var v37 uint8 = 1
        _ = fi35
        _ = v36
        _ = v37
    }
    v34 = v34
    var v38 uint64 = 18446744073709551615
    v5 -= 173
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v5
}
