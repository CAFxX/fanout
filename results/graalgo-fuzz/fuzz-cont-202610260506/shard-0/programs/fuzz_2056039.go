package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0(p0 int, p1 int8, p2 int) bool {
    switch (-(p0)) {
    case -136:
        for fi5 := 0; fi5 < 6; fi5++ {
            p0 -= (p2 ^ (p2 - fi5))
            _ = fi5
        }
    }
    p1 = (p1 & 127)
    if false {
        var v6 bool = true
        _ = v6
    }
    v7 := GenS1{F0: 90, F1: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}
func genFn1(p0 int8) uint8 {
    p0 += 1
    v8 := GenS2{F0: 2147483647}
    var v9 float64 = (1e10 - -2.25)
    var v10 int16 = 0
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 68
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 1
    var v3 string = "hello"
    v4 := make(map[uint32]string)
    v2 -= (v2 | (^(v2)))
    for fi11 := 0; fi11 < 6; fi11++ {
        if (false && ("a" < v3)) {
            v3 = "a"
        } else {
            if true {
                fmt.Println(v2, v3, fi11)
            } else {
                for fi12 := 0; fi12 < 1; fi12++ {
                    fmt.Println(fi11)
                    var v13 uint16 = 0
                    _ = len(v1)
                    v3 = v3
                    fmt.Println(fi12, v2, fi11)
                    _ = fi12
                    _ = v13
                }
                v2 = 299
            }
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v3 = "hello"
        v2 = 2
        if false {
            var v15 uint = 42
            v16 := GenS2{F0: (^(v2))}
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    var v17 uint = 1
    v18 := v1[(v3 + "")]
    fmt.Println(v2, v2)
    if (false || true) {
        var v19 uint32 = v2
        var v20 int = 410
        _ = v19
        _ = v20
    }
    switch "a" {
    case "hello":
        if true {
            v3 = ("" + (((v3 + v3) + v3) + "\n\t"))
            var v21 uint64 = 4294967295
            if ((v17 ^ 2) <= (v17 / ((255) | 1))) {
                v17 -= (v17 >> ((v17) & 63))
            }
            _ = v21
        } else {
            v22 := GenS2{F0: (v18 | (v2 << (((^(v17))) & 63)))}
            var v23 bool = (32767 > -32768)
            fmt.Println(v3, v2, v3)
            _ = v22
            _ = v23
        }
    }
    if false {
        var v24 uint32 = (v2 + (v18 / (((^(v18))) | 1)))
        _ = v24
    } else {
        if false {
            v3 = ("\n\t" + v3)
            v18 = v2
            fmt.Println(v2, v18)
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := []int{fi25, fi25, 4294967295}
            _ = fi25
            _ = v26
        }
    }
    var v27 int32 = 0
    for fi28 := 0; fi28 < 19; fi28++ {
        var v29 int16 = 42
        switch "hello" {
        case "foo bar":
            v30 := GenS3{F0: -94, F1: genFn1(70)}
            _ = v30
        }
        _ = fi28
        _ = v29
    }
    var v31 uint8 = 62
    var v32 uint64 = 42
    var v33 int32 = (v27 % ((-2147483648) | 1))
    var v34 float64 = ((1.5 + -2.25) / float64(-2.25))
    v17 += v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
