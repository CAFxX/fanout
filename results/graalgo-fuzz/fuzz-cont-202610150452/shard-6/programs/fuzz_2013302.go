package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
func genFn0() uint64 {
    var v4 uint8 = 1
    if false {
        var v5 string = ("" + "\n\t")
        for fi6 := 0; fi6 < 9; fi6++ {
            fmt.Println(v4, fi6, fi6)
            fmt.Println(v5, v5)
            fmt.Println(v4, fi6)
            _ = fi6
        }
        v5 = v5
        _ = v5
    }
    _ = v4
    return 479
}
func genFn1() int64 {
    for fi7 := 0; fi7 < 7; fi7++ {
        switch "a" {
        case "\n\t":
            var v8 int = (fi7 | (^(fi7)))
            _ = v8
        case "foo bar":
            var v9 int16 = 1000
            _ = v9
        case "hello":
            var v10 uint16 = 0
            _ = v10
        default:
            var v11 uint16 = 255
            _ = v11
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 int8 = 42
            v14 := []int{(fi12 % ((fi12) | 1)), (-(fi7))}
            _ = fi12
            _ = v13
            _ = v14
        }
        switch (fi7 / ((42) | 1)) {
        case -1:
            fmt.Println(fi7)
        default:
            var v15 int8 = -1
            _ = v15
        }
        _ = fi7
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        switch "\n\t" {
        case "\n\t":
            var v17 int32 = 1
            _ = v17
        }
        var v18 bool = true
        _ = fi16
        _ = v18
    }
    if false {
        switch ("foo bar" + "") {
        case "\n\t":
            switch ("\n\t" + "a") {
            case "foo bar":
                var v19 uint32 = 42
                _ = v19
            case "hello":
                if (!(!(167 < 0))) {
                    if false {
                        var v20 bool = true
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 2; fi21++ {
                            var v22 int32 = -602
                            v22 = (^(v22))
                            _ = fi21
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 13; fi23++ {
                            v24 := []int{(-(fi23)), (fi23 >> ((255) & 63)), fi23, -748}
                            _ = fi23
                            _ = v24
                        }
                    }
                    if (false || true) {
                        v25 := []int{869, 42, 42}
                        var v26 int16 = -1
                        if false {
                            for fi27 := 0; fi27 < 16; fi27++ {
                                var v28 string = ("\n\t" + "hello")
                                _ = fi27
                                _ = v28
                            }
                        }
                        _ = v25
                        _ = v26
                    }
                }
            case "\n\t":
                var v29 int8 = 63
                _ = v29
            }
        }
    } else {
        var v30 int16 = 854
        fmt.Println(v30)
        _ = v30
    }
    if ((!false) && ((1e10 - 0.1) == float64((1.5 + 0.1)))) {
        var v31 uint16 = 255
        fmt.Println(v31)
        _ = v31
    } else {
        if false {
            v32 := []int{-9223372036854775808, 128}
            _ = v32
        }
    }
    return 89
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 1
    var v3 uint16 = 441
    fmt.Println(v3)
    v33 := GenS1{F0: (!true)}
    v34 := GenS1{F0: true}
    v35 := GenS1{F0: (255 >= 2)}
    if false {
        v3 = v3
    } else {
        var v36 uint16 = (v3 & v3)
        if (78 > -1) {
            var v37 uint32 = 2147483648
            _ = v37
        } else {
            v36 += 540
        }
        _ = v36
    }
    var v38 int16 = -90
    fmt.Println(v1, v38, v2)
    for fi39 := 0; fi39 < 20; fi39++ {
        if true {
            if ((65535 > 2) && true) {
                v40 := []int{fi39}
                var v41 int64 = 0
                for fi42 := 0; fi42 < 20; fi42++ {
                    v2 -= (v2 - (v2 << ((18446744073709551615) & 63)))
                    _ = fi42
                }
                _ = v40
                _ = v41
            }
            if ((float64((-2.25 - 3.14159)) == -2.25) || false) {
                fmt.Println(v38)
                switch (("hello" + (("a" + (("foo bar" + "") + "")) + ("\n\t" + ("" + "")))) + "foo bar") {
                case "foo bar":
                    fmt.Println(v3)
                case "hello":
                    var v43 uint64 = 18446744073709551615
                    _ = v43
                default:
                    v44 := GenS1{F0: true}
                    _ = v44
                }
                v45 := GenS2{F0: genFn0()}
                _ = v45
            }
        }
        for fi46 := 0; fi46 < 11; fi46++ {
            if (!(-1 > -9223372036854775808)) {
                v47 := []int{(fi39 | -734), fi39}
                var v48 uint32 = 65535
                _ = v47
                _ = v48
            }
            _ = fi46
        }
        var v49 uint = 2
        _ = fi39
        _ = v49
    }
    v50 := []int{1}
    _ = len(v50)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v50
}
