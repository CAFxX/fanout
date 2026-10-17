package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int32; F1 int; F2 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 uint32 }
func genFn0(p0 bool, p1 uint16) uint32 {
    for fi4 := 0; fi4 < 16; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            if p0 {
                var v6 int32 = -1
                var v7 uint = 42
                var v8 uint64 = 42
                _ = v6
                _ = v7
                _ = v8
            } else {
                var v9 uint64 = 4294967295
                fmt.Println(fi4)
                if ((3.14159 * 0.0) != (-2.25 * (3.14159 * float64(1.5)))) {
                    v9 = 85
                    var v10 int16 = 0
                    _ = v10
                } else {
                    v11 := GenS1{F0: p0}
                    var v12 bool = p0
                    var v13 int = fi4
                    _ = v11
                    _ = v12
                    _ = v13
                }
                _ = v9
            }
            var v14 int16 = -1
            var v15 uint32 = 42
            _ = fi5
            _ = v14
            _ = v15
        }
        if p0 {
            switch 9223372036854775807 {
            case 32767:
                p0 = true
            case 437:
                for fi16 := 0; fi16 < 2; fi16++ {
                    switch "hello" {
                    case "\n\t":
                        if false {
                            p0 = p0
                            p0 = (0 < 9)
                        }
                    case "foo bar":
                        v17 := []int{}
                        _ = v17
                    }
                    _ = fi16
                }
            default:
                var v18 uint = 65535
                _ = v18
            }
            var v19 int8 = -1
            _ = v19
        } else {
            v20 := GenS2{F0: 2}
            _ = v20
        }
        _ = fi4
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        p1 = 404
        switch (-(fi21)) {
        case -621:
            if p0 {
                v22 := GenS3{F0: 100, F1: false}
                _ = v22
            } else {
                switch (fi21 << ((121) & 63)) {
                case 32768:
                    var v23 int16 = 1000
                    _ = v23
                case 960:
                    fmt.Println(p1, fi21, p0)
                default:
                    p0 = p0
                }
                fmt.Println(p0, p1, p1)
            }
        }
        _ = fi21
    }
    if p0 {
        for fi24 := 0; fi24 < 12; fi24++ {
            if p0 {
                if p0 {
                    v25 := GenS4{F0: 0, F1: (-(fi24)), F2: 500}
                    p1 = 0
                    _ = v25
                } else {
                    if (2 < 67) {
                        var v26 string = ""
                        if false {
                            var v27 int8 = -128
                            if p0 {
                                if p0 {
                                    for fi28 := 0; fi28 < 5; fi28++ {
                                        fmt.Println(p0, fi28)
                                        _ = fi28
                                    }
                                }
                                var v29 uint32 = 1
                                switch (v26 + v26) {
                                case "":
                                    fmt.Println(v26, v26)
                                case "foo bar":
                                    v30 := []int{fi24, 42}
                                    _ = v30
                                case "\n\t":
                                    var v31 uint8 = 127
                                    _ = v31
                                }
                                _ = v29
                            }
                            _ = v27
                        } else {
                            var v32 bool = (-35 <= 5)
                            _ = v32
                        }
                        _ = v26
                    }
                    var v33 string = "hello"
                    _ = v33
                }
                var v34 uint32 = 661
                v35 := GenS5{F0: (p1 & p1)}
                _ = v34
                _ = v35
            }
            for fi36 := 0; fi36 < 6; fi36++ {
                var v37 uint64 = 770
                v38 := GenS6{F0: 23, F1: uint32(v37)}
                switch "foo bar" {
                case "\n\t":
                    var v39 float64 = ((3.14159 * float64(0.0)) - -2.25)
                    _ = v39
                }
                _ = fi36
                _ = v37
                _ = v38
            }
            if (42 > 65535) {
                v40 := []int{}
                p1 = p1
                if len(v40) > 0 {
                    v40[0] = -116
                }
                _ = v40
            }
            _ = fi24
        }
    }
    if true {
        v41 := []int{}
        p0 = p0
        if (true && false) {
            if (255 != 42) {
                p0 = ("\n\t" != "")
            }
            p0 = (!p0)
        }
        _ = v41
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 528
    var v3 uint32 = 4294967295
    v3 = 964
    if (!true) {
        for fi42 := 0; fi42 < 1; fi42++ {
            if (42 != 2) {
                for fi43 := 0; fi43 < 6; fi43++ {
                    if true {
                        v2 -= v2
                        var v44 float64 = 3.14159
                        _ = v44
                    } else {
                        v2 = v2
                        if false {
                            var v45 uint8 = 255
                            var v46 uint32 = uint32(v2)
                            v45 = 0
                            _ = v45
                            _ = v46
                        }
                    }
                    var v47 float32 = 0.1
                    _ = fi43
                    _ = v47
                }
            } else {
                if (127 <= -1) {
                    var v48 int64 = 9223372036854775807
                    fmt.Println(v48)
                    _ = v48
                } else {
                    v49 := []int{(fi42 << ((0) & 63)), fi42, (fi42 | (fi42 - (^(fi42))))}
                    v50 := GenS5{F0: (v2 << ((18446744073709551615) & 63))}
                    _ = v49
                    _ = v50
                }
            }
            v51 := v1[v3]
            _ = fi42
            _ = v51
        }
    }
    switch ("a" + ("\n\t" + ("" + "foo bar"))) {
    case "hello":
        v52 := []int{2147483647, 206, 32768, 21}
        _ = v52
    }
    v3 += v3
    var v53 int64 = 250
    fmt.Println(v2, v3)
    fmt.Println(v3, v3)
    delete(v1, 2147483648)
    var v54 uint32 = 16
    v3 = (^(v54))
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
    _ = v54
}
