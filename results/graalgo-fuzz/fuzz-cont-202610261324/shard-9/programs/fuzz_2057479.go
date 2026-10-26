package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint8) uint8 {
    for fi7 := 0; fi7 < 9; fi7++ {
        if true {
            if (-24 < 0) {
                var v8 float32 = 1.5
                var v9 int8 = -22
                for fi10 := 0; fi10 < 11; fi10++ {
                    v11 := GenS1{F0: uint(p0)}
                    v12 := GenS1{F0: 42}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = v8
                _ = v9
            }
            v13 := []int{32767}
            _ = v13
        } else {
            var v14 uint8 = (^(p0))
            var v15 uint = 770
            _ = v14
            _ = v15
        }
        _ = fi7
    }
    v16 := GenS1{F0: 255}
    var v17 int32 = 1
    var v18 int64 = -889
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    return p0
}

func main() {
    var v1 uint16 = 32767
    var v2 int32 = 1
    v3 := []uint32{260}
    var v4 uint16 = (^(v1))
    v5 := make(map[string]bool)
    var v6 float64 = (((1e10 / 1e10) - 1.5) * ((float64((1e10 + 0.0)) * 0.0) - ((0.0 * (3.14159 + (0.1 * 3.14159))) + (0.1 * (0.1 + 3.14159)))))
    switch 128 {
    case 786:
        delete(v5, "a")
    }
    fmt.Println(v6, v2)
    v19 := v5["\n\t"]
    switch "" {
    case "foo bar":
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS1{F0: 0}
            var v22 uint8 = 128
            var v23 int32 = (v2 ^ v2)
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
    case "":
        if (!(("" < ("\n\t" + (("hello" + "a") + ""))) || (false || (((1 == 2147483647) && (18446744073709551615 > 730)) && (2 <= 0))))) {
            for fi24 := 0; fi24 < 3; fi24++ {
                fmt.Println(v6, v19)
                _ = len(v5)
                var v25 float64 = v6
                _ = fi24
                _ = v25
            }
            v4 = (v1 << ((4294967295) & 63))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            var v26 int32 = -1
            if v19 {
                fmt.Println(v26)
            }
            for fi27 := 0; fi27 < 5; fi27++ {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                v2 = -1
                _ = fi27
            }
            _ = v26
        }
    case "hello":
        _ = len(v5)
    }
    switch "" {
    case "":
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 float32 = 0.0
            var v30 uint64 = 42
            _ = fi28
            _ = v29
            _ = v30
        }
    case "\n\t":
        v6 += v6
    }
    v31 := GenS2{F0: v1, F1: v2}
    if v19 {
        var v32 int16 = 32767
        var v33 int8 = -119
        v34 := GenS3{F0: 18446744073709551615}
        _ = v32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        if v19 {
            if (false || (v19 && ((42 != 80) || (0 == 94)))) {
                var v36 int64 = -1
                for fi37 := 0; fi37 < 19; fi37++ {
                    v38 := GenS1{F0: 374}
                    _ = fi37
                    _ = v38
                }
                _ = v36
            }
            var v39 string = ("hello" + ("a" + ("a" + "a")))
            var v40 int32 = v2
            _ = v39
            _ = v40
        } else {
            var v41 string = ("foo bar" + "hello")
            _ = v41
        }
        for fi42 := 0; fi42 < 15; fi42++ {
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
            if (!v19) {
                v4 = 583
                v43 := GenS4{F0: 32767}
                v19 = false
                _ = v43
            }
            var v44 int16 = 32767
            _ = fi42
            _ = v44
        }
        for fi45 := 0; fi45 < 12; fi45++ {
            v46 := GenS5{F0: 0, F1: v1}
            v47 := GenS6{F0: 256}
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = fi35
    }
    v4 -= (v4 << ((0) & 63))
    v48 := GenS3{F0: 0}
    for fi49 := 0; fi49 < 8; fi49++ {
        delete(v5, ("a" + ("\n\t" + ("hello" + "hello"))))
        v2 = v2
        _ = fi49
    }
    v6 = v6
    var v50 uint8 = 128
    v50 = (v50 & 42)
    v50 -= (v50 % (((v50 >> ((611) & 63))) | 1))
    var v51 uint32 = 2147483647
    var v52 uint64 = 4294967295
    v2 = (-(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v48
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
