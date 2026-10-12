package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
func genFn0(p0 uint16, p1 int16, p2 int) int32 {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int64 = -1
        var v9 int64 = int64(p1)
        for fi10 := 0; fi10 < 12; fi10++ {
            switch (fi7 - fi7) {
            case -674:
                v11 := []int{fi10, (-(p2)), (fi10 / ((fi7) | 1))}
                _ = v11
            case 128:
                var v12 float32 = -2.25
                _ = v12
            default:
                v9 = (-(v9))
            }
            _ = fi10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v13 int32 = 2147483647
    if (83 < 103) {
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := []int{}
            var v16 uint8 = 101
            var v17 uint8 = v16
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        v18 := []int{}
        _ = v18
    } else {
        p1 = (p1 * p1)
        p2 = (p2 % (((p2 * -9223372036854775808)) | 1))
        for fi19 := 0; fi19 < 12; fi19++ {
            for fi20 := 0; fi20 < 13; fi20++ {
                fmt.Println(p1, v13, p2)
                p0 = p0
                var v21 uint = 4294967295
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    }
    fmt.Println(v13, p2)
    if false {
        var v22 uint16 = 82
        _ = v22
    } else {
        v23 := []int{p2, (p2 - (p2 + -530))}
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return (^(v13))
}

func main() {
    var v1 uint16 = 65535
    var v2 int = -988
    var v3 bool = (!(false || true))
    var v4 uint32 = 2147483647
    var v5 bool = (42 <= 0)
    var v6 uint32 = (^(v4))
    if (959 != v1) {
        v5 = v3
    } else {
        v24 := GenS1{F0: 738}
        _ = v24
    }
    v6 -= (^(v6))
    switch "" {
    case "a":
        var v25 uint = 148
        _ = v25
    }
    v6 -= v6
    var v26 float64 = (0.1 / ((float64((float32((0.1 + 1e10)) - (-2.25 * float32(0.1)))) + 3.14159) / 1e10))
    for fi27 := 0; fi27 < 13; fi27++ {
        if false {
            var v28 int = (v2 - (fi27 & 127))
            v29 := GenS2{F0: -9223372036854775808}
            _ = v28
            _ = v29
        }
        _ = fi27
    }
    var v30 uint16 = v1
    for fi31 := 0; fi31 < 16; fi31++ {
        v32 := GenS3{F0: ("\n\t" + "\n\t")}
        var v33 uint64 = 2
        var v34 int = v2
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    fmt.Println(v2, v26)
    if false {
        switch "\n\t" {
        case "a":
            var v35 uint32 = 2147483647
            _ = v35
        case "":
            v36 := []int{v2}
            _ = v36
        }
    }
    switch (v2 / ((v2) | 1)) {
    case 9223372036854775807:
        for fi37 := 0; fi37 < 8; fi37++ {
            for fi38 := 0; fi38 < 16; fi38++ {
                v6 -= 0
                v39 := []int{(fi38 * (v2 >> ((952) & 63)))}
                var v40 bool = v5
                _ = len(v39)
                _ = fi38
                _ = v39
                _ = v40
            }
            for fi41 := 0; fi41 < 19; fi41++ {
                switch "hello" {
                case "":
                    var v42 int32 = -2147483648
                    _ = v42
                default:
                    fmt.Println(v6)
                }
                fmt.Println(v2, v4)
                _ = fi41
            }
            _ = fi37
        }
    case 0:
        var v43 float32 = 1.5
        _ = v43
    case 72:
        var v44 int16 = -225
        _ = v44
    default:
        var v45 string = ("a" + "a")
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
