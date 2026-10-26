package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int16; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 int; F2 int }
func genFn0(p0 int32, p1 uint8) uint8 {
    if true {
        for fi4 := 0; fi4 < 5; fi4++ {
            if true {
                if true {
                    var v5 uint16 = 2
                    p0 = p0
                    _ = v5
                } else {
                    p1 += 2
                }
                fmt.Println(fi4, fi4, fi4)
                switch "a" {
                case "hello":
                    var v6 uint = 255
                    _ = v6
                case "a":
                    if true {
                        v7 := GenS1{F0: 189}
                        if false {
                            p1 = 38
                        }
                        for fi8 := 0; fi8 < 16; fi8++ {
                            if true {
                                p0 = (p0 / ((p0) | 1))
                                fmt.Println(fi8, p0, p0)
                                if (p0 <= p0) {
                                    fmt.Println(p1, p0)
                                    fmt.Println(fi8, p0)
                                    p1 = (^(p1))
                                } else {
                                    var v9 uint16 = 42
                                    if ((p0 % ((p0) | 1)) <= 1) {
                                        var v10 int32 = (p0 / (((p0 % (((p0 * p0)) | 1))) | 1))
                                        p1 = (^(p1))
                                        v11 := []int{(fi4 % (((fi8 << ((2) & 63))) | 1))}
                                        fmt.Println(v10, v9)
                                        _ = v10
                                        _ = v11
                                    }
                                    _ = v9
                                }
                            }
                            p1 = (p1 << ((255) & 63))
                            p0 = p0
                            var v12 string = "\n\t"
                            _ = fi8
                            _ = v12
                        }
                        _ = v7
                    }
                default:
                    v13 := GenS2{F0: (^(p1)), F1: 1, F2: (!false)}
                    _ = v13
                }
            }
            p0 = -659
            _ = fi4
        }
        v14 := GenS2{F0: p1, F1: 32767, F2: (2 > 387)}
        _ = v14
    } else {
        v15 := []int{185, 172}
        var v16 uint = 127
        _ = v15
        _ = v16
    }
    var v17 string = "foo bar"
    var v18 float32 = 0.1
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return (^(p1))
}
func genFn1() uint16 {
    var v19 string = "hello"
    v20 := []int{}
    _ = v19
    _ = v20
    return 2
}

func main() {
    v1 := []uint32{214, 42, 2147483647, 2}
    v2 := make(map[string]string)
    var v3 float32 = -2.25
    var v21 uint = 255
    v22 := GenS2{F0: genFn0(2147483647, 42), F1: -1, F2: false}
    if false {
        if len(v1) > 0 {
            v1[0] = 1
        }
        var v23 int64 = -866
        _ = v23
    } else {
        var v24 uint16 = genFn1()
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        v3 = float32((0.1 * float64(v3)))
        _ = fi25
    }
    if ((!true) && (1 == 42)) {
        var v26 uint8 = 127
        if (false && false) {
            v3 -= 0.1
            for fi27 := 0; fi27 < 20; fi27++ {
                for fi28 := 0; fi28 < 20; fi28++ {
                    var v29 uint = (v21 >> (((v21 - (v21 ^ 4294967295))) & 63))
                    _ = len(v2)
                    if ((v26 | 0) == 255) {
                        switch 1 {
                        case -1:
                            switch (fi28 - fi28) {
                            case -2147483648:
                                _ = len(v1)
                            }
                        case 2147483647:
                            switch "\n\t" {
                            case "":
                                v30 := GenS3{F0: fi27}
                                _ = v30
                            }
                        case 538:
                            var v31 uint = 1
                            _ = v31
                        default:
                            var v32 int32 = -1
                            _ = v32
                        }
                    }
                    _ = fi28
                    _ = v29
                }
                var v33 uint16 = 2
                _ = fi27
                _ = v33
            }
        }
        if (!true) {
            var v34 int = 659
            var v35 string = ((("foo bar" + "hello") + "a") + (("foo bar" + "") + "foo bar"))
            _ = v34
            _ = v35
        } else {
            for fi36 := 0; fi36 < 12; fi36++ {
                fmt.Println(v26)
                _ = fi36
            }
            v3 = -2.25
        }
        _ = v26
    }
    v37 := GenS1{F0: 888}
    v1 = append(v1, 701)
    for fi38 := 0; fi38 < 3; fi38++ {
        var v39 string = "a"
        if false {
            fmt.Println(fi38)
            fmt.Println(v3)
            var v40 uint8 = 245
            _ = v40
        }
        _ = len(v2)
        _ = fi38
        _ = v39
    }
    if ((!(2147483647 == 1)) || true) {
        if (!(1e10 < (1e10 / (-2.25 / 3.14159)))) {
            var v41 int = 42
            for fi42 := 0; fi42 < 16; fi42++ {
                v43 := GenS2{F0: 127, F1: 1000, F2: (42 == 9223372036854775807)}
                _ = fi42
                _ = v43
            }
            _ = v41
        } else {
            var v44 int8 = 100
            if true {
                var v45 int16 = 32767
                fmt.Println(v44)
                v46 := GenS4{F0: v45, F1: -20, F2: -2147483648}
                _ = v45
                _ = v46
            }
            _ = v44
        }
    }
    switch ("\n\t" + "") {
    case "":
        if len(v1) > 0 {
            v1[0] = 1
        }
    }
    var v47 uint64 = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v47
}
