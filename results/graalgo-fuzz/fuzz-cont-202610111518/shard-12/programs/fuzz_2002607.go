package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 bool; F2 bool }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint32; F1 uint64; F2 int8 }
func genFn0() uint8 {
    var v5 uint8 = 0
    v6 := GenS1{F0: 1}
    fmt.Println(v5)
    for fi7 := 0; fi7 < 16; fi7++ {
        switch ("hello" + "") {
        case "\n\t":
            v8 := []int{fi7, (fi7 / ((fi7) | 1)), int(v5)}
            _ = v8
        default:
            v9 := GenS2{F0: 747, F1: (3.14159 <= 1.5), F2: true}
            _ = v9
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 int32, p1 uint32, p2 int32) int16 {
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 int = (fi10 & (fi11 | fi10))
            var v13 int16 = 42
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 1000
}

func main() {
    var v1 int = 300
    v2 := make(map[string]int)
    var v3 string = ("" + ("a" + ("" + ("" + ("foo bar" + "a")))))
    var v4 int8 = 42
    if (127 > 1) {
        v14 := GenS1{F0: 0}
        var v15 uint = 339
        _ = v14
        _ = v15
    } else {
        v16 := GenS3{F0: (^(v4)), F1: ""}
        _ = v16
    }
    v1 = 1
    if (v1 >= 602) {
        var v17 int = 32768
        v18 := []int{(v1 >> ((255) & 63)), v1}
        _ = v17
        _ = v18
    }
    v2[v3] = (v1 >> ((65535) & 63))
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS2{F0: 846, F1: (false && (false || (!true))), F2: true}
        _ = fi19
        _ = v20
    }
    switch "\n\t" {
    case "a":
        v21 := GenS2{F0: 4294967295, F1: (true || (false && (!(!(186 == 166))))), F2: (!(true && (genFn1(1, 2147483648, -2147483648) < 32767)))}
        _ = v21
    case "hello":
        var v22 int = v1
        _ = v22
    case "\n\t":
        v23 := v2[(v3 + ("a" + (v3 + (v3 + ("hello" + (((v3 + "") + "foo bar") + v3))))))]
        _ = v23
    }
    switch (v3 + "hello") {
    case "foo bar":
        for fi24 := 0; fi24 < 2; fi24++ {
            if ((64 == 89) && (true || false)) {
                var v25 uint32 = 585
                v3 = v3
                var v26 int = (-(v1))
                _ = v25
                _ = v26
            } else {
                fmt.Println(v3)
                v4 = v4
                switch (fi24 >> ((905) & 63)) {
                case -9223372036854775808:
                    for fi27 := 0; fi27 < 19; fi27++ {
                        v1 += 128
                        var v28 string = v3
                        var v29 int16 = -1
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                case 2147483647:
                    v3 = (v3 + (("" + "") + ""))
                case 42:
                    v30 := v2["foo bar"]
                    _ = v30
                }
            }
            _ = fi24
        }
    case "":
        for fi31 := 0; fi31 < 11; fi31++ {
            v3 = v3
            v3 = v3
            for fi32 := 0; fi32 < 8; fi32++ {
                v1 = (fi31 & -1)
                fmt.Println(v3, v3)
                _ = fi32
            }
            _ = fi31
        }
    default:
        for fi33 := 0; fi33 < 14; fi33++ {
            v34 := []int{(v1 % ((v1) | 1)), 936}
            _ = fi33
            _ = v34
        }
    }
    switch v1 {
    case 128:
        if (1 >= 18446744073709551615) {
            v1 = (v1 - 127)
        } else {
            var v35 string = "hello"
            v3 = "a"
            _ = v35
        }
    case 32768:
        var v36 int32 = 2147483647
        _ = v36
    }
    v3 = ("hello" + ((v3 + "") + v3))
    v1 = v1
    v37 := v2[(v3 + v3)]
    if (!(18446744073709551615 < 526)) {
        fmt.Println(v1, v3, v37)
        v4 -= -84
        switch 0 {
        case -9223372036854775808:
            for fi38 := 0; fi38 < 4; fi38++ {
                var v39 int = -327
                var v40 uint32 = 42
                _ = fi38
                _ = v39
                _ = v40
            }
        case 127:
            for fi41 := 0; fi41 < 9; fi41++ {
                for fi42 := 0; fi42 < 8; fi42++ {
                    var v43 int8 = -111
                    var v44 int = -536
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                v45 := GenS4{F0: 65535, F1: 0, F2: 100}
                _ = fi41
                _ = v45
            }
        case 957:
            v3 = ((v3 + "a") + (v3 + ""))
        default:
            v46 := []int{42, (v1 % ((v37) | 1)), (v1 ^ (^(v1))), -165}
            _ = v46
        }
    }
    v3 = "\n\t"
    for fi47 := 0; fi47 < 2; fi47++ {
        var v48 bool = false
        var v49 int32 = 2147483647
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
