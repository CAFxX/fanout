package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
func genFn0(p0 int16) int {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{}
        var v8 int32 = 2147483647
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint32 = 719
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := []int{fi11, (fi11 | fi11)}
        p0 = (p0 + 0)
        var v13 float64 = 3.14159
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{42, -600, -9223372036854775808, 42}
    var v15 uint16 = 32767
    _ = p0
    _ = v14
    _ = v15
    return 32768
}
func genFn1() int16 {
    var v16 uint = 0
    v16 += (v16 % ((v16) | 1))
    v17 := []int{9223372036854775807}
    v17 = append(v17, -240)
    _ = v16
    _ = v17
    return 32767
}

func main() {
    var v1 bool = true
    var v2 int32 = 302
    v3 := []uint32{65535, 145}
    var v4 int16 = 0
    var v5 uint = 0
    v1 = v1
    if ((!v1) || ((-(v2)) <= 2147483647)) {
        switch "hello" {
        case "":
            v5 = v5
        case "\n\t":
            if v1 {
                var v18 int8 = 42
                _ = v18
            } else {
                v4 = genFn1()
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        default:
            v19 := GenS1{F0: (v2 & 2147483647)}
            _ = v19
        }
        var v20 uint = (v5 + v5)
        _ = v20
    }
    fmt.Println(v1, v2)
    var v21 int = genFn0((v4 >> ((v5) & 63)))
    var v22 string = "\n\t"
    switch v22 {
    case "":
        v23 := GenS2{F0: (!(true || v1))}
        _ = v23
    case "foo bar":
        for fi24 := 0; fi24 < 6; fi24++ {
            var v25 int32 = (v2 + (v2 & v2))
            switch v22 {
            case "hello":
                for fi26 := 0; fi26 < 5; fi26++ {
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    var v27 bool = v1
                    v4 += 1
                    _ = fi26
                    _ = v27
                }
            case "a":
                fmt.Println(v5)
            default:
                switch (fi24 % ((1) | 1)) {
                case 1:
                    var v28 int32 = 1
                    _ = v28
                }
            }
            v29 := GenS2{F0: (v22 <= ((v22 + v22) + (("foo bar" + ((("hello" + ("\n\t" + v22)) + "a") + v22)) + "hello")))}
            _ = fi24
            _ = v25
            _ = v29
        }
    case "\n\t":
        var v30 uint8 = 1
        _ = v30
    default:
        var v31 int32 = (v2 >> ((v5) & 63))
        _ = v31
    }
    if false {
        v4 -= v4
    } else {
        v5 -= 794
        var v32 bool = v1
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 uint16 = 368
            _ = fi33
            _ = v34
        }
        _ = v32
    }
    v2 -= v2
    for fi35 := 0; fi35 < 8; fi35++ {
        _ = len(v3)
        _ = fi35
    }
    v1 = (428 == v5)
    for fi36 := 0; fi36 < 3; fi36++ {
        _ = len(v3)
        var v37 uint16 = 65535
        v4 = 42
        _ = fi36
        _ = v37
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
