package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int16, p1 int8) uint {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint16 = 32767
        _ = fi6
        _ = v7
    }
    p0 -= p0
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 bool, p1 uint8, p2 uint32) int32 {
    p2 += (p2 * (p2 | (^(p2))))
    p0 = false
    p2 += p2
    v8 := []int{42, int(p2), -1, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2147483647
}

func main() {
    var v1 int8 = 127
    var v2 int = 418
    v3 := []string{((("\n\t" + "") + "a") + "foo bar")}
    var v4 bool = (false && false)
    var v5 uint32 = 2147483647
    v2 += (v2 | (v2 & (^(v2))))
    v9 := GenS1{F0: 2, F1: 1, F2: 42}
    var v10 int = 2147483647
    for fi11 := 0; fi11 < 4; fi11++ {
        switch "a" {
        case "foo bar":
            var v12 float32 = (0.1 - 3.14159)
            _ = v12
        case "a":
            if true {
                v2 = 713
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        var v13 uint64 = 193
        _ = fi11
        _ = v13
    }
    fmt.Println(v1, v2)
    switch 0 {
    case -1:
        var v14 bool = (587 != 40)
        _ = v14
    }
    v1 = 42
    if v4 {
        var v15 int32 = genFn1((42 != 1), 42, v5)
        _ = v15
    }
    switch -371 {
    case -16:
        v5 += v5
    case 850:
        switch 438 {
        case -9223372036854775808:
            v1 += (v1 % ((25) | 1))
        case -274:
            v16 := GenS1{F0: 746, F1: -629, F2: 490}
            _ = v16
        }
    default:
        switch -1 {
        case -2147483648:
            if false {
                v1 += (v1 + v1)
                v17 := GenS2{F0: 42}
                _ = v17
            }
        case 32767:
            v5 += v5
        case 4294967295:
            if (!(!(2 > 42))) {
                fmt.Println(v1, v4, v1)
            }
        }
    }
    fmt.Println(v2, v4)
    v18 := GenS1{F0: 65535, F1: -1, F2: 42}
    for fi19 := 0; fi19 < 15; fi19++ {
        v4 = (42 <= 42)
        v20 := GenS2{F0: 28}
        _ = fi19
        _ = v20
    }
    v21 := GenS3{F0: 0}
    switch ("\n\t" + "\n\t") {
    case "hello":
        var v22 uint8 = 1
        _ = v22
    default:
        for fi23 := 0; fi23 < 10; fi23++ {
            for fi24 := 0; fi24 < 14; fi24++ {
                switch (^(v2)) {
                case 2147483647:
                    var v25 string = (("hello" + "") + "foo bar")
                    _ = v25
                case 9223372036854775807:
                    v3 = append(v3, "a")
                }
                v2 += 772
                _ = fi24
            }
            _ = fi23
        }
    }
    v5 = (v5 / ((v5) | 1))
    var v26 int = v10
    if true {
        v1 += v1
    } else {
        switch v2 {
        case 0:
            var v27 int8 = v1
            _ = v27
        case -253:
            var v28 uint16 = 32767
            _ = v28
        case 32767:
            v29 := GenS1{F0: 2, F1: 2147483647, F2: 0}
            _ = v29
        }
        v10 = v26
        var v30 string = ("\n\t" + "hello")
        _ = v30
    }
    fmt.Println(v1, v2, v26)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
