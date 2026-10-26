package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint64, p1 int32) uint32 {
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 uint64 = (p0 & (p0 >> ((1) & 63)))
            var v9 uint64 = (v8 >> ((42) & 63))
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v10 := GenS1{F0: 100, F1: -74, F2: -9223372036854775808}
    p1 += p1
    switch "" {
    case "":
        fmt.Println(p1)
    case "foo bar":
        for fi11 := 0; fi11 < 8; fi11++ {
            p0 += p0
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = v10
    return 858
}

func main() {
    var v1 uint = 2
    var v2 int64 = -530
    var v3 int8 = 1
    var v4 int64 = v2
    var v5 float64 = (float64(0.0) * 1e10)
    var v12 uint8 = 2
    var v13 int8 = v3
    v13 = v13
    var v14 int8 = (-(v3))
    switch -9223372036854775808 {
    case 32767:
        switch -720 {
        case -9223372036854775808:
            var v15 float64 = 0.0
            _ = v15
        case 127:
            for fi16 := 0; fi16 < 15; fi16++ {
                if true {
                    v2 = v4
                    v13 += v13
                } else {
                    var v17 bool = (v1 > (^(v1)))
                    _ = v17
                }
                v3 = v14
                switch (fi16 << ((v1) & 63)) {
                case -900:
                    v18 := []int{(-(fi16)), (-(fi16))}
                    _ = v18
                case -629:
                    switch ("hello" + "foo bar") {
                    case "\n\t":
                        var v19 int = fi16
                        _ = v19
                    case "foo bar":
                        if (9223372036854775807 <= -1) {
                            for fi20 := 0; fi20 < 10; fi20++ {
                                v1 -= 42
                                var v21 int = (-(fi16))
                                var v22 float64 = (0.0 + (1e10 * -2.25))
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                            fmt.Println(v5)
                        }
                    }
                default:
                    v13 += -79
                }
                _ = fi16
            }
        case 2147483647:
            v23 := []int{128, 9223372036854775807, 182}
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 5; fi25++ {
            var v26 bool = (0 >= -32768)
            fmt.Println(fi25, v13, v1)
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    fmt.Println(v12)
    if (true && (!true)) {
        switch 32768 {
        case -1:
            v27 := GenS2{F0: ""}
            _ = v27
        case -9223372036854775808:
            if ((v5 * 3.14159) <= v5) {
                v3 = -1
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    v12 -= 2
                }
            }
        }
        if (false || false) {
            switch "foo bar" {
            case "":
                if ((v1 / ((uint(v12)) | 1)) <= v1) {
                    for fi28 := 0; fi28 < 17; fi28++ {
                        v13 = v3
                        if (0 < 695) {
                            v29 := []int{fi28, (fi28 + (fi28 - (fi28 ^ fi28)))}
                            _ = v29
                        }
                        var v30 string = ("foo bar" + "hello")
                        _ = fi28
                        _ = v30
                    }
                }
            }
            v31 := GenS2{F0: ""}
            _ = v31
        } else {
            v32 := []int{766}
            _ = v32
        }
        v14 = (^(v3))
    }
    switch 4294967295 {
    case -885:
        var v33 bool = false
        _ = v33
    case 32768:
        switch -9223372036854775808 {
        case 9223372036854775807:
            v34 := []int{}
            _ = v34
        }
    default:
        if (v14 == (^(v14))) {
            v35 := GenS3{F0: 42}
            _ = v35
        }
    }
    var v36 int = 806
    var v37 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
