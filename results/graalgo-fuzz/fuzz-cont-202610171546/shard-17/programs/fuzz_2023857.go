package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int, p1 int8, p2 uint64) uint32 {
    fmt.Println(p0)
    var v4 uint16 = 84
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1(p0 int8) int {
    p0 -= (p0 << ((18446744073709551615) & 63))
    if (((!false) && (4294967295 > 0)) || (3.14159 < 1e10)) {
        for fi5 := 0; fi5 < 6; fi5++ {
            switch int(p0) {
            case 827:
                p0 = p0
            case 127:
                v6 := GenS1{F0: 42}
                _ = v6
            case 2147483647:
                v7 := []int{fi5}
                _ = v7
            }
            var v8 string = "\n\t"
            _ = fi5
            _ = v8
        }
        v9 := GenS1{F0: 2}
        if true {
            for fi10 := 0; fi10 < 20; fi10++ {
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := []int{}
                    _ = fi11
                    _ = v12
                }
                var v13 uint8 = 128
                _ = fi10
                _ = v13
            }
            var v14 int64 = -1
            _ = v14
        }
        _ = v9
    }
    _ = p0
    return 128
}

func main() {
    var v1 float64 = 1e10
    var v2 int = 9223372036854775807
    v3 := []string{"\n\t"}
    var v15 int32 = -566
    if false {
        for fi16 := 0; fi16 < 11; fi16++ {
            for fi17 := 0; fi17 < 12; fi17++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v1 -= ((3.14159 * (1.5 - -2.25)) - (v1 / v1))
                fmt.Println(v1, fi16)
                _ = fi17
            }
            v1 = ((v1 * 0.0) + -2.25)
            var v18 int32 = v15
            _ = fi16
            _ = v18
        }
        fmt.Println(v2, v15, v15)
        var v19 uint64 = 1
        _ = v19
    } else {
        v20 := GenS1{F0: 165}
        _ = v20
    }
    switch "\n\t" {
    case "\n\t":
        v21 := GenS1{F0: 42}
        _ = v21
    case "hello":
        v15 -= -1
    case "":
        v15 = v15
    default:
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 int16 = 493
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 uint = 134
                _ = fi24
                _ = v25
            }
            _ = fi22
            _ = v23
        }
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        for fi27 := 0; fi27 < 14; fi27++ {
            v28 := GenS1{F0: 32767}
            fmt.Println(v2, v2)
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 7; fi29++ {
            if false {
                v1 = 1.5
            }
            v1 = v1
            var v30 int = genFn1(100)
            _ = fi29
            _ = v30
        }
        _ = fi26
    }
    v15 = v15
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 float64 = v1
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        if false {
            v2 = 4294967295
            _ = len(v3)
            var v34 float64 = (float64(-2.25) - (((v1 - float64(1e10)) - 0.0) * ((v1 + v1) - v1)))
            _ = v34
        }
        _ = fi33
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    fmt.Println(v1, v2, v1)
    for fi35 := 0; fi35 < 19; fi35++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v3 = append(v3, ("a" + "a"))
        _ = fi35
    }
    var v36 uint8 = 0
    v1 -= v1
    var v37 uint16 = 1
    for fi38 := 0; fi38 < 4; fi38++ {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = fi38
    }
    v36 = v36
    fmt.Println(v2)
    if false {
        switch "a" {
        case "hello":
            v1 += v1
        }
        v37 = (v37 + v37)
        v15 = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}
