package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8, p1 int, p2 uint16) int {
    var v6 int8 = (p0 ^ (-(p0)))
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{}
        p2 = p2
        v6 += (^(v6))
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "hello":
        var v9 string = "hello"
        _ = v9
    case "a":
        var v10 uint32 = 4294967295
        _ = v10
    case "foo bar":
        var v11 uint16 = p2
        _ = v11
    default:
        v12 := []int{-9223372036854775808, (p1 / ((p1) | 1)), p1, (-(p1))}
        _ = v12
    }
    if true {
        for fi13 := 0; fi13 < 17; fi13++ {
            fmt.Println(p2, fi13, fi13)
            _ = fi13
        }
        var v14 uint = 802
        _ = v14
    }
    p1 = (p1 << ((0) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p1 + 42)
}

func main() {
    var v1 float32 = -2.25
    var v2 int8 = 42
    var v3 uint8 = 255
    var v4 int16 = 1
    var v5 int8 = (^(v2))
    if (uint64(v3) < 18446744073709551615) {
        if ((!false) || false) {
            var v15 int32 = 42
            _ = v15
        } else {
            v2 = v5
        }
        if (!true) {
            switch 128 {
            case -1:
                var v16 int64 = -467
                _ = v16
            case -615:
                fmt.Println(v4, v4, v3)
            }
            v4 = v4
            v3 = v3
        } else {
            v17 := []int{2147483647, -166, 32767, 4294967295}
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 uint = 4294967295
                v1 = 1e10
                var v20 int = (fi18 / ((-1) | 1))
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 bool = (18446744073709551615 <= 609)
            _ = v17
            _ = v21
        }
    } else {
        v22 := []int{-1, 32767, 32768, 0}
        _ = v22
    }
    var v23 float32 = v1
    var v24 uint8 = 42
    switch int(v4) {
    case 42:
        v2 = (v2 & (v2 & (v5 | (v5 & v5))))
    }
    v25 := []int{1}
    if (true && (true || ((-2.25 * 0.0) <= (1.5 * (1e10 / 0.1))))) {
        if false {
            for fi26 := 0; fi26 < 6; fi26++ {
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                v23 += v23
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                _ = fi26
            }
            switch "" {
            case "foo bar":
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
            case "a":
                v5 -= v5
            case "\n\t":
                v27 := GenS1{F0: (v24 << ((307) & 63)), F1: 1, F2: 642}
                _ = v27
            }
        } else {
            v28 := GenS2{F0: 255}
            var v29 int = -9223372036854775808
            fmt.Println(v2)
            _ = v28
            _ = v29
        }
        if ((682 == 1) || ((2 < 32767) && true)) {
            v24 = v3
        } else {
            var v30 int64 = -614
            if false {
                _ = len(v25)
                v5 -= v2
            } else {
                var v31 float32 = v1
                _ = v31
            }
            _ = v30
        }
        for fi32 := 0; fi32 < 4; fi32++ {
            fmt.Println(v5)
            for fi33 := 0; fi33 < 9; fi33++ {
                var v34 uint8 = (v24 | 255)
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
    } else {
        var v35 int8 = v5
        _ = v35
    }
    fmt.Println(v24, v3, v5)
    for fi36 := 0; fi36 < 2; fi36++ {
        switch (((("foo bar" + ((("hello" + "hello") + "\n\t") + "\n\t")) + "\n\t") + (("a" + ("hello" + (("foo bar" + "\n\t") + "a"))) + ("" + "\n\t"))) + "a") {
        case "\n\t":
            v37 := GenS2{F0: 1}
            _ = v37
        default:
            for fi38 := 0; fi38 < 5; fi38++ {
                if ((fi36 << ((42) & 63)) != fi36) {
                    if (((-2.25 + 3.14159) + 3.14159) > 0.0) {
                        _ = len(v25)
                    } else {
                        if len(v25) > 0 {
                            v25[0] = (fi36 & (fi38 | fi36))
                        }
                        var v39 int = (fi36 + 128)
                        var v40 uint16 = 42
                        _ = v39
                        _ = v40
                    }
                    v24 = v24
                    v1 = (v23 + (v23 * 1.5))
                    var v41 uint64 = 18446744073709551615
                    if (9223372036854775807 != 42) {
                        fmt.Println(v23, v2, fi38)
                        v41 = v41
                    }
                    _ = v41
                }
                _ = fi38
            }
        }
        v2 = 100
        v42 := GenS3{F0: 1}
        _ = fi36
        _ = v42
    }
    v2 += (v5 & 100)
    v4 -= v4
    if (-1 <= 2147483647) {
        if false {
            if len(v25) > 0 {
                v25[0] = 32768
            }
            v43 := GenS1{F0: 203, F1: 186, F2: 42}
            _ = v43
        } else {
            var v44 float32 = ((v1 * v23) + (v23 - 0.1))
            _ = v44
        }
    }
    fmt.Println(v3)
    v4 = v4
    var v45 string = ((("" + "\n\t") + "") + ((("hello" + "a") + ("\n\t" + "hello")) + "\n\t"))
    for fi46 := 0; fi46 < 17; fi46++ {
        var v47 uint = 18446744073709551615
        v45 = "hello"
        for fi48 := 0; fi48 < 20; fi48++ {
            var v49 int16 = -32768
            var v50 int8 = v5
            if true {
                v45 = ""
                var v51 uint8 = (v3 >> ((v47) & 63))
                _ = v51
            }
            _ = fi48
            _ = v49
            _ = v50
        }
        _ = fi46
        _ = v47
    }
    v5 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
