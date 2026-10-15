package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int64, p1 string) uint8 {
    v4 := []int{-1, 9223372036854775807, 42}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi5 := 0; fi5 < 1; fi5++ {
        p1 = "foo bar"
        switch fi5 {
        case 1:
            var v6 int8 = -91
            _ = v6
        case -9223372036854775808:
            p1 = p1
        default:
            var v7 int = fi5
            _ = v7
        }
        v8 := GenS1{F0: p1}
        _ = fi5
        _ = v8
    }
    p1 = p1
    var v9 string = (p1 + "")
    _ = p0
    _ = p1
    _ = v4
    _ = v9
    return 1
}
func genFn1(p0 uint, p1 int64) uint16 {
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: "\n\t"}
    if (!false) {
        var v13 int32 = -1
        fmt.Println(v13, v13, p1)
        if true {
            v13 = v13
            v14 := GenS1{F0: "hello"}
            _ = v14
        }
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v12
    return 718
}

func main() {
    v1 := []uint32{1, 2147483648}
    var v2 int32 = 1
    v3 := []int{50, -705}
    v2 += (v2 << ((18446744073709551615) & 63))
    v2 = v2
    var v15 string = ("hello" + ("hello" + ("\n\t" + (("\n\t" + "\n\t") + "hello"))))
    if true {
        v16 := GenS1{F0: v15}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            v2 += (-(v2))
            fmt.Println(v2)
            fmt.Println(fi17, v15, v2)
            _ = fi17
        }
    }
    fmt.Println(v2, v15)
    var v18 uint32 = 601
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            if (1 == 829) {
                var v21 string = v15
                var v22 uint64 = 42
                v22 += (v22 + (^(v22)))
                _ = v21
                _ = v22
            }
            v18 += v18
            var v23 uint8 = 90
            _ = fi20
            _ = v23
        }
        var v24 float32 = (((1.5 * ((0.1 + (3.14159 * ((1e10 * float32(3.14159)) - 0.1))) / 1.0)) / -2.25) + float32(1.5))
        _ = fi19
        _ = v24
    }
    v3 = append(v3, -81)
    var v25 uint = 2
    switch 0 {
    case 2147483647:
        for fi26 := 0; fi26 < 6; fi26++ {
            _ = len(v1)
            var v27 uint = 255
            _ = fi26
            _ = v27
        }
    case 1:
        v2 = (-(v2))
    case 0:
        v3 = append(v3, int(v18))
    }
    v3 = append(v3, 4294967295)
    if (false && ((140 > 2) && ((4294967295 != 1) || true))) {
        for fi28 := 0; fi28 < 10; fi28++ {
            if false {
                switch v15 {
                case "foo bar":
                    v25 = v25
                }
                var v29 int16 = 32767
                _ = v29
            }
            var v30 uint8 = 255
            switch (^(fi28)) {
            case 128:
                v2 = (v2 | v2)
            case -2147483648:
                fmt.Println(fi28, v25)
            case 4294967295:
                v2 += (-(v2))
            }
            _ = fi28
            _ = v30
        }
        v25 += 42
    } else {
        v15 = (v15 + "hello")
        v18 += (v18 >> ((65535) & 63))
        if (v2 > v2) {
            var v31 uint64 = uint64(v18)
            v15 = "a"
            v32 := GenS1{F0: ""}
            _ = v31
            _ = v32
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v25
    _ = v3
}
