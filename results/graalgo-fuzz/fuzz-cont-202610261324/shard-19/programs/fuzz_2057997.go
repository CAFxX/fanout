package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 int }
func genFn0(p0 bool) uint32 {
    v6 := GenS1{F0: 42, F1: "foo bar", F2: 42}
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: fi8, F1: "\n\t", F2: 255}
            var v10 uint32 = 65535
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 uint8 = 128
        _ = fi7
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := []int{fi12}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v6
    return 642
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = -2.25
    var v3 uint16 = 153
    v4 := []uint32{683, 42, 4294967295}
    v5 := []int{646}
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int64 = 552
        if false {
            v16 := GenS1{F0: (fi14 - (fi14 | int(v3))), F1: (("" + ((("" + ("foo bar" + ("foo bar" + ("hello" + "")))) + ("foo bar" + "hello")) + "")) + ("a" + ("\n\t" + "hello"))), F2: 225}
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    var v17 int8 = -1
    switch "hello" {
    case "a":
        v2 = v2
    }
    if (!(32767 < -1)) {
        v17 = -1
        v3 -= (v3 << ((669) & 63))
    }
    fmt.Println(v17)
    for fi18 := 0; fi18 < 2; fi18++ {
        if (42 != 1) {
            v19 := GenS2{F0: 42}
            var v20 uint32 = genFn0(true)
            _ = v19
            _ = v20
        }
        _ = fi18
    }
    v21 := GenS3{F0: 1}
    if (2 != genFn0(true)) {
        if true {
            var v22 int = 1
            switch 658 {
            case -2147483648:
                v5 = append(v5, (v22 % ((32767) | 1)))
            case 32767:
                _ = len(v5)
            case 2147483647:
                fmt.Println(v3, v17)
            }
            _ = v22
        } else {
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 uint = 65535
                if false {
                    var v25 bool = false
                    if v25 {
                        var v26 uint8 = 128
                        _ = v26
                    }
                    var v27 int64 = 0
                    _ = v25
                    _ = v27
                } else {
                    if (true && false) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        fmt.Println(v1, v24)
                    } else {
                        var v28 float64 = float64((-2.25 * float32((1.5 + v2))))
                        _ = v28
                    }
                    if true {
                        for fi29 := 0; fi29 < 20; fi29++ {
                            if (true && (42 > 42)) {
                                v30 := GenS4{F0: (fi23 % ((fi23) | 1))}
                                _ = v30
                            } else {
                                v31 := GenS4{F0: 32767}
                                v32 := GenS5{F0: -32768, F1: fi29}
                                _ = v31
                                _ = v32
                            }
                            var v33 uint = (v24 >> ((42) & 63))
                            var v34 int64 = 739
                            _ = fi29
                            _ = v33
                            _ = v34
                        }
                        var v35 int = int(v24)
                        _ = v35
                    }
                }
                _ = fi23
                _ = v24
            }
            v3 = v3
        }
        for fi36 := 0; fi36 < 7; fi36++ {
            fmt.Println(v17, v3, v17)
            if len(v5) > 0 {
                v5[0] = (fi36 - 4294967295)
            }
            _ = fi36
        }
        if ("hello" < "foo bar") {
            v1 += (v1 << ((0) & 63))
        }
    } else {
        fmt.Println(v17, v3)
        for fi37 := 0; fi37 < 8; fi37++ {
            switch "hello" {
            case "hello":
                v38 := GenS2{F0: -1}
                _ = v38
            case "a":
                for fi39 := 0; fi39 < 2; fi39++ {
                    var v40 uint16 = 2
                    _ = fi39
                    _ = v40
                }
            }
            _ = fi37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
