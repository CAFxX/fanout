package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 bool; F2 int16 }

func main() {
    var v1 int16 = 42
    v2 := make(map[int]uint32)
    var v3 uint = 18446744073709551615
    var v4 uint32 = 2147483647
    var v5 bool = false
    var v6 int = 2147483647
    var v7 uint16 = 2
    var v8 int8 = 0
    var v9 int64 = -9223372036854775808
    switch (v6 >> ((v3) & 63)) {
    case 510:
        var v10 uint64 = 25
        _ = v10
    }
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v7 = uint16(v3)
    default:
        v11 := GenS1{F0: 127, F1: (-(v9))}
        _ = v11
    }
    var v12 uint16 = (v7 ^ 32767)
    if v5 {
        if (!false) {
            var v13 string = ("foo bar" + "hello")
            v5 = true
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        switch v6 {
        case -394:
            var v15 bool = v5
            _ = v15
        case 4294967295:
            var v16 int64 = 1
            _ = v16
        case 260:
            if (((1 >= -1) || ((31 == 42) && true)) || ((v4 ^ (v4 / ((0) | 1))) == (^(v4)))) {
                v5 = ((v3 << ((v3) & 63)) != v3)
                v17 := GenS2{F0: ("\n\t" + "foo bar")}
                _ = v17
            } else {
                var v18 int8 = -120
                _ = v18
            }
        default:
            var v19 uint32 = v4
            _ = v19
        }
        var v20 float64 = float64((3.14159 * float32(1e10)))
        if v5 {
            switch "a" {
            case "hello":
                delete(v2, (v6 + (fi14 % (((-(v6))) | 1))))
            case "":
                var v21 float32 = 3.14159
                _ = v21
            }
            v6 = fi14
        }
        _ = fi14
        _ = v20
    }
    var v22 float64 = 1e10
    for fi23 := 0; fi23 < 15; fi23++ {
        if v5 {
            for fi24 := 0; fi24 < 3; fi24++ {
                v1 = (v1 ^ 1000)
                v25 := []int{(-(v6)), (fi23 / ((v6) | 1))}
                _ = fi24
                _ = v25
            }
        }
        _ = fi23
    }
    v26 := []int{}
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    _ = len(v26)
    v27 := GenS1{F0: (v6 | (v6 + (v6 % ((v6) | 1)))), F1: 42}
    for fi28 := 0; fi28 < 14; fi28++ {
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
        for fi29 := 0; fi29 < 12; fi29++ {
            v7 -= v7
            v30 := GenS1{F0: fi29, F1: (v9 / ((v9) | 1))}
            var v31 uint32 = v4
            _ = fi29
            _ = v30
            _ = v31
        }
        switch int(v3) {
        case -213:
            v9 -= 197
        case 0:
            if false {
                v8 += (-(v8))
                var v32 float32 = (-2.25 / ((0.1 / 1.0) - 3.14159))
                for fi33 := 0; fi33 < 5; fi33++ {
                    switch -441 {
                    case 39:
                        _ = len(v2)
                    default:
                        var v34 string = ("\n\t" + ("\n\t" + ("a" + ("foo bar" + ""))))
                        if true {
                            var v35 uint32 = v4
                            v36 := GenS3{F0: (v1 * v1)}
                            _ = v35
                            _ = v36
                        } else {
                            fmt.Println(v1)
                            fmt.Println(fi28, fi28)
                        }
                        _ = v34
                    }
                    v37 := GenS2{F0: "\n\t"}
                    _ = fi33
                    _ = v37
                }
                _ = v32
            } else {
                fmt.Println(v22, v3, v5)
                v5 = (fi28 >= v6)
                if v5 {
                    if len(v26) > 0 {
                        v26[0] = (fi28 << (((v3 / ((v3) | 1))) & 63))
                    }
                    var v38 bool = v5
                    _ = v38
                } else {
                    fmt.Println(v4, v8, v1)
                }
            }
        }
        _ = fi28
    }
    v6 += 32767
    if v5 {
        if v5 {
            var v39 float64 = float64(float32(v22))
            v40 := GenS2{F0: ("hello" + "a")}
            _ = v39
            _ = v40
        }
        switch ("hello" + "a") {
        case "a":
            v9 = (^(v9))
        case "hello":
            var v41 float32 = 1e10
            _ = v41
        default:
            if (2147483647 >= 763) {
                if (v3 <= (v3 << ((uint(v4)) & 63))) {
                    fmt.Println(v8)
                    v5 = true
                    if (248 == 208) {
                        v42 := GenS3{F0: (-(v1))}
                        var v43 int32 = 42
                        switch (("foo bar" + "foo bar") + "hello") {
                        case "a":
                            if false {
                                var v44 uint8 = 255
                                v45 := GenS4{F0: 0, F1: true, F2: -1}
                                var v46 bool = true
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                        case "foo bar":
                            if (uint8(v4) >= 2) {
                                for fi47 := 0; fi47 < 12; fi47++ {
                                    fmt.Println(v5, fi47, v22)
                                    v43 = v43
                                    v48 := v2[(-(v6))]
                                    _ = fi47
                                    _ = v48
                                }
                                for fi49 := 0; fi49 < 2; fi49++ {
                                    for fi50 := 0; fi50 < 20; fi50++ {
                                        v51 := v2[(v6 & fi50)]
                                        if v5 {
                                            v5 = v5
                                            fmt.Println(v6, v51, v7)
                                            var v52 uint32 = (^(v51))
                                            v53 := GenS3{F0: (v1 ^ -33)}
                                            _ = v52
                                            _ = v53
                                        }
                                        _ = fi50
                                        _ = v51
                                    }
                                    _ = fi49
                                }
                            }
                        }
                        _ = v42
                        _ = v43
                    }
                }
                v3 = v3
            } else {
                var v54 uint16 = v7
                fmt.Println(v22, v5)
                _ = v54
            }
        }
    } else {
        delete(v2, (-(v6)))
    }
    delete(v2, (v6 * (v6 & (v6 >> ((505) & 63)))))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
