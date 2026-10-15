package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 int; F1 bool }
func genFn0(p0 int64, p1 int8) uint16 {
    var v7 int16 = -76
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            v7 += v7
            _ = fi9
        }
        _ = fi8
    }
    var v10 int64 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 1
}

func main() {
    var v1 float32 = 3.14159
    var v2 int64 = 9223372036854775807
    var v3 int64 = (v2 | (v2 + v2))
    var v4 uint32 = 1
    var v5 uint = 19
    v6 := make(map[uint32]bool)
    var v11 string = "\n\t"
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 int64 = v3
        var v14 bool = (!true)
        v15 := v6[v4]
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        v3 = v2
        switch (fi16 / ((fi16) | 1)) {
        case -431:
            if (574 < 18446744073709551615) {
                for fi17 := 0; fi17 < 16; fi17++ {
                    v18 := GenS1{F0: uint8(v4), F1: v5, F2: 42}
                    _ = fi17
                    _ = v18
                }
                v19 := []int{42}
                _ = v19
            } else {
                v20 := GenS1{F0: 1, F1: 4294967295, F2: 1}
                _ = v20
            }
        case 2147483647:
            _ = len(v6)
        }
        var v21 uint = (v5 & v5)
        _ = fi16
        _ = v21
    }
    if true {
        v22 := v6[(^(v4))]
        var v23 float64 = 1e10
        _ = v22
        _ = v23
    }
    if false {
        for fi24 := 0; fi24 < 13; fi24++ {
            v1 = v1
            var v25 uint = 255
            _ = fi24
            _ = v25
        }
        switch ((v11 + "") + "a") {
        case "hello":
            delete(v6, (v4 >> (((v5 | v5)) & 63)))
        }
    } else {
        v2 += v3
        for fi26 := 0; fi26 < 6; fi26++ {
            delete(v6, 2)
            var v27 float32 = 1.5
            v27 += v1
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            v29 := GenS1{F0: uint8(v5), F1: v5, F2: 2147483647}
            var v30 uint32 = (v4 | uint32(v5))
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    if (v5 < v5) {
        for fi31 := 0; fi31 < 4; fi31++ {
            v11 = "\n\t"
            _ = fi31
        }
    }
    var v32 uint = 1
    if (!("" < v11)) {
        if false {
            v5 -= (v5 ^ (v32 - v5))
            v33 := GenS2{F0: 4294967295, F1: false}
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 int8 = 100
                _ = fi34
                _ = v35
            }
            _ = v33
        } else {
            var v36 float32 = v1
            for fi37 := 0; fi37 < 12; fi37++ {
                v38 := []int{}
                if (!true) {
                    v4 += uint32(v5)
                    v6[uint32(v32)] = (-983 > (fi37 | (fi37 << ((42) & 63))))
                    v3 = (v2 << (((^(v32))) & 63))
                } else {
                    v4 -= (v4 & (v4 ^ (v4 ^ v4)))
                }
                _ = fi37
                _ = v38
            }
            _ = v36
        }
        if (!false) {
            for fi39 := 0; fi39 < 6; fi39++ {
                v11 = (("\n\t" + "foo bar") + "\n\t")
                _ = fi39
            }
            fmt.Println(v5)
        } else {
            for fi40 := 0; fi40 < 4; fi40++ {
                for fi41 := 0; fi41 < 17; fi41++ {
                    v42 := GenS1{F0: 2, F1: v32, F2: -2147483648}
                    _ = fi41
                    _ = v42
                }
                _ = fi40
            }
            if false {
                if false {
                    v5 = 255
                }
            }
        }
        fmt.Println(v4, v2)
    }
    if false {
        var v43 uint = v32
        v43 = v32
        if true {
            if (149 >= 200) {
                var v44 uint = 255
                _ = v44
            } else {
                var v45 float64 = -2.25
                for fi46 := 0; fi46 < 5; fi46++ {
                    for fi47 := 0; fi47 < 5; fi47++ {
                        v3 = 1
                        var v48 int16 = -517
                        _ = fi47
                        _ = v48
                    }
                    _ = fi46
                }
                fmt.Println(v4)
                _ = v45
            }
        }
        _ = v43
    } else {
        var v49 int8 = -25
        var v50 string = "a"
        _ = v49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
