package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16) uint {
    v6 := []int{576}
    p0 -= (p0 << ((2) & 63))
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int8 = 9
        v9 := GenS1{F0: fi7}
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 float64 = 1.5
            var v12 int16 = 0
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if (32767 <= 32767) {
        switch "foo bar" {
        case "a":
            if (1 != -9223372036854775808) {
                v13 := GenS1{F0: -9223372036854775808}
                _ = v13
            } else {
                var v14 int64 = 546
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v14
            }
        case "hello":
            var v15 uint64 = 1
            _ = v15
        case "foo bar":
            if true {
                fmt.Println(p0)
                var v16 int32 = 42
                _ = v16
            }
        default:
            var v17 int16 = 992
            _ = v17
        }
        fmt.Println(p0)
        fmt.Println(p0)
    } else {
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 uint8 = 0
            switch ("hello" + ("a" + "a")) {
            case "\n\t":
                v6 = append(v6, (fi18 + fi18))
            case "hello":
                v20 := GenS1{F0: fi18}
                _ = v20
            case "foo bar":
                for fi21 := 0; fi21 < 11; fi21++ {
                    fmt.Println(fi21, fi21)
                    switch (fi21 / ((fi18) | 1)) {
                    case 9223372036854775807:
                        var v22 uint = 504
                        _ = v22
                    }
                    _ = len(v6)
                    _ = fi21
                }
            default:
                var v23 bool = (!(!false))
                _ = v23
            }
            v19 += v19
            _ = fi18
            _ = v19
        }
        fmt.Println(p0)
    }
    _ = p0
    _ = v6
    return 255
}
func genFn1(p0 int16, p1 string) int32 {
    var v24 int = -1
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 uint = 2
        _ = fi25
        _ = v26
    }
    var v27 int64 = 703
    v28 := GenS1{F0: v24}
    v29 := []int{}
    _ = p0
    _ = p1
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    return -1
}

func main() {
    var v1 int32 = 2147483647
    v2 := []uint32{286}
    var v3 int64 = -9223372036854775808
    var v4 float32 = ((0.0 / 0.1) * 0.0)
    v5 := make(map[int]bool)
    var v30 string = ("a" + "foo bar")
    var v31 uint8 = 176
    for fi32 := 0; fi32 < 18; fi32++ {
        switch v30 {
        case "foo bar":
            v33 := GenS2{F0: 0}
            _ = v33
        }
        _ = fi32
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        v30 = v30
        _ = fi34
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 int = 32768
        var v37 int64 = v3
        var v38 uint32 = 0
        _ = fi35
        _ = v36
        _ = v37
        _ = v38
    }
    var v39 uint64 = 227
    v4 -= (-2.25 + ((1e10 - 0.0) + v4))
    v40 := GenS2{F0: 243}
    switch -971 {
    case 32768:
        v2 = append(v2, 546)
    case 0:
        var v41 uint16 = 65535
        _ = v41
    case -189:
        v5[127] = (v31 <= (v31 >> ((981) & 63)))
    default:
        switch ("foo bar" + "\n\t") {
        case "a":
            v4 = (v4 / ((v4 + (-2.25 / (float32(0.1) + (v4 - v4)))) * (float32(3.14159) + (1.5 * (3.14159 * v4)))))
        }
    }
    switch v30 {
    case "":
        for fi42 := 0; fi42 < 17; fi42++ {
            fmt.Println(fi42, v31)
            var v43 int64 = 156
            _ = fi42
            _ = v43
        }
    case "hello":
        var v44 int32 = (v1 * v1)
        _ = v44
    default:
        var v45 int8 = 0
        _ = v45
    }
    var v46 int8 = 40
    var v47 bool = true
    var v48 float32 = (v4 * 1.5)
    var v49 int64 = (v3 % ((-118) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}
