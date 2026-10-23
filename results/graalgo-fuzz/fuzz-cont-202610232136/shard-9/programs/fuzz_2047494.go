package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint8, p1 int8) uint16 {
    v7 := []int{989, 4294967295, 128, 736}
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 int8 = (p1 / ((-128) | 1))
        switch (fi8 >> ((2) & 63)) {
        case -2147483648:
            var v10 int32 = 589
            _ = v10
        case 372:
            var v11 float64 = 3.14159
            _ = v11
        default:
            fmt.Println(p1, p0, p1)
        }
        _ = fi8
        _ = v9
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        switch "" {
        case "":
            v13 := GenS1{F0: (fi12 - 127)}
            _ = v13
        }
        if len(v7) > 0 {
            v7[0] = fi12
        }
        _ = fi12
    }
    var v14 float32 = float32((0.1 / -2.25))
    _ = p0
    _ = p1
    _ = v14
    _ = v7
    return 42
}
func genFn1(p0 int32, p1 uint64, p2 int16) uint8 {
    if true {
        for fi15 := 0; fi15 < 17; fi15++ {
            p2 = p2
            switch (fi15 * fi15) {
            case 42:
                var v16 int16 = p2
                _ = v16
            default:
                if (!false) {
                    var v17 int16 = -1
                    var v18 uint8 = 255
                    _ = v17
                    _ = v18
                }
            }
            _ = fi15
        }
        switch "foo bar" {
        case "hello":
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 float64 = 0.1
                _ = fi19
                _ = v20
            }
        case "":
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 uint64 = 639
                p2 = (p2 >> ((426) & 63))
                v23 := GenS1{F0: fi21}
                _ = fi21
                _ = v22
                _ = v23
            }
        default:
            for fi24 := 0; fi24 < 11; fi24++ {
                v25 := GenS1{F0: fi24}
                fmt.Println(fi24)
                _ = fi24
                _ = v25
            }
        }
        var v26 float64 = 0.0
        _ = v26
    } else {
        var v27 float64 = (float64(-2.25) / (float64(0.1) * float64((0.0 + 0.1))))
        p0 = p0
        v28 := []int{42}
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 int16 = p2
        p0 -= (p0 % ((339) | 1))
        fmt.Println(p1)
        _ = fi29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    var v1 uint8 = 207
    var v2 float64 = float64(1e10)
    var v3 float32 = 1e10
    var v4 uint8 = v1
    var v5 int16 = 1000
    v6 := make(map[string]string)
    if (!(532 > 2147483648)) {
        for fi31 := 0; fi31 < 4; fi31++ {
            if true {
                for fi32 := 0; fi32 < 5; fi32++ {
                    var v33 int64 = 42
                    switch (fi31 << ((1) & 63)) {
                    case -9223372036854775808:
                        var v34 int = (fi31 >> ((18446744073709551615) & 63))
                        _ = v34
                    }
                    _ = fi32
                    _ = v33
                }
            }
            _ = fi31
        }
    }
    var v35 bool = true
    var v36 uint = 18446744073709551615
    v37 := GenS1{F0: -640}
    v1 += v4
    var v38 int32 = 269
    var v39 float64 = (v2 + float64(1.5))
    var v40 int = -361
    switch "" {
    case "a":
        if (!(!(0 <= 42))) {
            switch (v40 << ((v36) & 63)) {
            case 580:
                v2 = 1.5
            case -482:
                v5 -= (v5 - v5)
            }
        } else {
            v41 := []int{v40, 0, -9223372036854775808}
            fmt.Println(v36, v36, v1)
            _ = v41
        }
    default:
        var v42 int64 = 9223372036854775807
        _ = v42
    }
    v43 := []int{(v40 & (v40 % ((32767) | 1))), -1}
    var v44 uint32 = 65535
    var v45 int8 = 127
    switch v40 {
    case -1:
        v46 := GenS1{F0: (v40 ^ (v40 & v40))}
        _ = v46
    case -9223372036854775808:
        v6[""] = "foo bar"
    case 32768:
        delete(v6, "hello")
    default:
        for fi47 := 0; fi47 < 19; fi47++ {
            var v48 uint32 = (v44 % ((v44) | 1))
            switch ("hello" + ("a" + (("hello" + "") + "foo bar"))) {
            case "a":
                delete(v6, "\n\t")
            }
            fmt.Println(v36, v3)
            _ = fi47
            _ = v48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
