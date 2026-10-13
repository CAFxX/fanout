package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint; F1 int8; F2 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS6 struct { F0 bool }
func genFn0() int32 {
    v4 := GenS1{F0: 254}
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint64 = 0
        if (0 <= 703) {
            v7 := GenS1{F0: 42}
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := []int{}
                v6 -= (v6 | v6)
                v6 = v6
                _ = fi8
                _ = v9
            }
            switch (fi5 % ((32768) | 1)) {
            case -400:
                for fi10 := 0; fi10 < 19; fi10++ {
                    v11 := GenS2{F0: 42, F1: 42}
                    v12 := []int{9223372036854775807, fi10, fi10, (fi10 << ((255) & 63))}
                    v13 := GenS1{F0: 0}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            case 1:
                v14 := GenS2{F0: 1, F1: 190}
                _ = v14
            case -9223372036854775808:
                var v15 float32 = 1e10
                _ = v15
            }
            _ = v7
        } else {
            var v16 uint16 = 2
            v17 := []int{}
            _ = v16
            _ = v17
        }
        _ = fi5
        _ = v6
    }
    var v18 int = 9223372036854775807
    v18 -= v18
    _ = v18
    _ = v4
    return -260
}
func genFn1(p0 int64, p1 int8) uint16 {
    if false {
        v19 := GenS1{F0: 255}
        for fi20 := 0; fi20 < 13; fi20++ {
            if (3.14159 != 1e10) {
                var v21 string = ("a" + ("foo bar" + ("hello" + ("foo bar" + "foo bar"))))
                _ = v21
            } else {
                for fi22 := 0; fi22 < 10; fi22++ {
                    if (2 < 164) {
                        var v23 float64 = 1e10
                        p1 = p1
                        _ = v23
                    }
                    _ = fi22
                }
                fmt.Println(p0, fi20)
                v24 := GenS1{F0: 189}
                _ = v24
            }
            if true {
                fmt.Println(p1, fi20, p1)
            } else {
                for fi25 := 0; fi25 < 6; fi25++ {
                    v26 := GenS1{F0: 2}
                    _ = fi25
                    _ = v26
                }
                p1 -= (p1 >> ((65535) & 63))
                var v27 uint = 2
                _ = v27
            }
            for fi28 := 0; fi28 < 20; fi28++ {
                var v29 int8 = p1
                _ = fi28
                _ = v29
            }
            _ = fi20
        }
        _ = v19
    } else {
        v30 := []int{540, int(p0), -1}
        for fi31 := 0; fi31 < 17; fi31++ {
            if false {
                if (761 == 65535) {
                    fmt.Println(fi31)
                    fmt.Println(fi31, p0)
                }
            }
            _ = fi31
        }
        p1 -= p1
        _ = v30
    }
    var v32 uint = 18446744073709551615
    v33 := GenS2{F0: v32, F1: 1000}
    var v34 uint32 = 0
    v32 = v32
    _ = p0
    _ = p1
    _ = v32
    _ = v33
    _ = v34
    return 153
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    if (!true) {
        var v35 int64 = -880
        _ = v35
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        v37 := GenS3{F0: 248, F1: -128, F2: true}
        if (!(4294967295 <= 42)) {
            var v38 int8 = 93
            _ = v38
        } else {
            for fi39 := 0; fi39 < 4; fi39++ {
                var v40 float64 = float64(-2.25)
                var v41 int = (fi39 % (((fi39 - (fi36 | 0))) | 1))
                _ = fi39
                _ = v40
                _ = v41
            }
            var v42 uint16 = 0
            _ = v42
        }
        _ = fi36
        _ = v37
    }
    switch ("\n\t" + "a") {
    case "foo bar":
        v43 := GenS4{F0: 0}
        _ = v43
    default:
        if (true || false) {
            var v44 int32 = 1
            fmt.Println(v44)
            var v45 int16 = -34
            _ = v44
            _ = v45
        }
    }
    v46 := v3[1]
    var v47 int32 = 42
    for fi48 := 0; fi48 < 17; fi48++ {
        var v49 uint64 = 1
        if false {
            v49 = (^(v49))
        } else {
            fmt.Println(v47, v47, v47)
            var v50 string = "a"
            v51 := GenS5{F0: 936, F1: 0, F2: 841}
            _ = v50
            _ = v51
        }
        v52 := GenS6{F0: true}
        _ = fi48
        _ = v49
        _ = v52
    }
    v46 -= (v46 - (v46 / ((0) | 1)))
    v2[("hello" + (("foo bar" + "\n\t") + "hello"))] = 394
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v47
}
