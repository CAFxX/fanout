package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int8; F1 int16 }

func main() {
    var v1 int32 = 42
    v2 := []string{(("foo bar" + "\n\t") + ("\n\t" + "a"))}
    var v3 int64 = 1
    var v4 uint64 = 1
    var v5 uint64 = v4
    if len(v2) > 0 {
        v2[0] = ("hello" + ("a" + ""))
    }
    if (98 >= 42) {
        var v6 uint64 = 1
        v7 := GenS1{F0: 2, F1: 42}
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: 261, F1: 32767}
            if len(v2) > 0 {
                v2[0] = ("\n\t" + ("" + "a"))
            }
            if true {
                var v10 string = "foo bar"
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 int16 = -247
                    v12 = 32767
                    var v13 uint8 = 203
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    } else {
        v5 = 472
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint8 = 1
            var v16 uint64 = 1
            _ = fi14
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 11; fi17++ {
            v18 := GenS1{F0: 2, F1: 255}
            v19 := GenS1{F0: 42, F1: 42}
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    if false {
        v2 = append(v2, "\n\t")
        if false {
            v1 = 925
        } else {
            var v20 int = 4294967295
            _ = v20
        }
    }
    v4 = v5
    switch (("foo bar" + "\n\t") + "foo bar") {
    case "a":
        for fi21 := 0; fi21 < 14; fi21++ {
            if true {
                var v22 int8 = 42
                switch ("\n\t" + "") {
                case "a":
                    v1 -= v1
                case "\n\t":
                    v23 := GenS2{F0: (^(v5)), F1: 42, F2: (^(v4))}
                    _ = v23
                case "":
                    v5 = 0
                }
                _ = v22
            }
            v24 := GenS3{F0: -1, F1: 1}
            var v25 uint32 = 256
            _ = fi21
            _ = v24
            _ = v25
        }
    }
    if false {
        v26 := GenS3{F0: 127, F1: 32767}
        _ = v26
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        v3 = (v3 >> ((1) & 63))
        v28 := GenS2{F0: (v5 | v5), F1: 65535, F2: 18446744073709551615}
        _ = fi27
        _ = v28
    }
    switch 42 {
    case 0:
        for fi29 := 0; fi29 < 17; fi29++ {
            v5 -= v5
            for fi30 := 0; fi30 < 14; fi30++ {
                var v31 string = (((("" + "hello") + ("a" + "a")) + ("a" + "a")) + "hello")
                v32 := GenS1{F0: 319, F1: 2}
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi29
        }
    default:
        v5 += (^(v5))
    }
    if (!true) {
        if (56 <= 4294967295) {
            var v33 int8 = 127
            var v34 int = -9223372036854775808
            _ = v33
            _ = v34
        }
        for fi35 := 0; fi35 < 1; fi35++ {
            for fi36 := 0; fi36 < 19; fi36++ {
                var v37 uint = 0
                _ = fi36
                _ = v37
            }
            for fi38 := 0; fi38 < 19; fi38++ {
                if false {
                    v39 := GenS1{F0: 0, F1: 212}
                    _ = v39
                } else {
                    v5 += (v5 * v4)
                    fmt.Println(fi38, fi38, fi38)
                }
                var v40 float64 = (3.14159 + -2.25)
                v41 := GenS2{F0: (v5 / ((v4) | 1)), F1: 65535, F2: v5}
                _ = fi38
                _ = v40
                _ = v41
            }
            _ = fi35
        }
    } else {
        v1 += (-(v1))
        var v42 int8 = 42
        if (!(v42 <= 127)) {
            v2 = append(v2, "foo bar")
        }
        _ = v42
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        v1 = 872
        v3 -= (v3 - (v3 / ((v3) | 1)))
        _ = fi43
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
