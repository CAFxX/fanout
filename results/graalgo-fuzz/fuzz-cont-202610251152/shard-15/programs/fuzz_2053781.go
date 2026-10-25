package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int }
type GenS2 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint64, p1 uint64, p2 int16) bool {
    v4 := GenS1{F0: 32767, F1: (p2 * p2), F2: 32767}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (false || ((((255 <= 460) || ("hello" <= "foo bar")) || (-2.25 == float32(float64(0.1)))) && true))
}
func genFn1(p0 uint16, p1 int, p2 int16) int16 {
    var v5 int32 = 602
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: 2, F1: (p2 | p2), F2: 613}
        for fi8 := 0; fi8 < 1; fi8++ {
            v5 = v5
            v9 := []int{-141}
            fmt.Println(p0, fi6)
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 int8 = 0
            if (true && (!false)) {
                var v12 float64 = -2.25
                _ = v12
            } else {
                v13 := []int{}
                v11 = -120
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        switch (p1 >> ((42) & 63)) {
        case 32767:
            v15 := GenS1{F0: (p0 - 32767), F1: p2, F2: 0}
            _ = v15
        }
        for fi16 := 0; fi16 < 5; fi16++ {
            v17 := []int{32767}
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            _ = fi16
            _ = v17
        }
        _ = fi14
    }
    v18 := GenS1{F0: 32767, F1: (p2 & (p2 << ((65535) & 63))), F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v5
    return (p2 ^ p2)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{}
    var v3 uint16 = 32767
    v3 = v3
    var v19 uint32 = 651
    for fi20 := 0; fi20 < 4; fi20++ {
        v3 = v3
        if (!(!true)) {
            fmt.Println(fi20, v3)
        }
        _ = fi20
    }
    if false {
        if true {
            for fi21 := 0; fi21 < 13; fi21++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if (0 > -1) {
                    switch ("hello" + "foo bar") {
                    case "hello":
                        v19 = v19
                    case "\n\t":
                        var v22 string = "foo bar"
                        _ = v22
                    }
                    v3 = v3
                }
                v19 += 542
                _ = fi21
            }
            switch "a" {
            case "hello":
                switch "hello" {
                case "a":
                    v23 := GenS1{F0: (^(v3)), F1: 32767, F2: 127}
                    _ = v23
                }
            default:
                delete(v1, "foo bar")
            }
        }
    }
    v24 := GenS1{F0: v3, F1: 638, F2: -350}
    for fi25 := 0; fi25 < 13; fi25++ {
        switch "" {
        case "a":
            v26 := GenS2{F0: 255, F1: (v3 - (^(v3)))}
            _ = v26
        default:
            fmt.Println(fi25, v19)
        }
        v3 += (^(v3))
        _ = fi25
    }
    v3 = (^(v3))
    var v27 uint = 490
    var v28 uint64 = 18446744073709551615
    if true {
        v28 = 2
    } else {
        v27 += 255
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
}
