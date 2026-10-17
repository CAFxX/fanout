package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint64, p1 uint64, p2 uint) int8 {
    fmt.Println(p0, p2, p1)
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 float64 = 0.1
        var v7 uint = p2
        var v8 uint32 = 4294967295
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    p1 -= (^(p1))
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1() uint64 {
    var v9 uint32 = 65535
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int = (fi10 >> ((65535) & 63))
        _ = fi10
        _ = v11
    }
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint32 = 2147483647
    var v2 float64 = (-2.25 / 0.1)
    var v3 uint32 = v1
    v4 := make(map[int]bool)
    delete(v4, 2147483647)
    v1 += 42
    if true {
        delete(v4, 627)
        delete(v4, 759)
        if ((true || (!(145 >= 226))) || false) {
            var v12 string = ""
            _ = v12
        } else {
            v3 = v3
        }
    } else {
        var v13 string = ("a" + ("\n\t" + "a"))
        fmt.Println(v2, v2, v1)
        var v14 uint16 = 32767
        _ = v13
        _ = v14
    }
    var v15 int = -151
    v16 := []int{(v15 ^ (-(v15))), (v15 / (((v15 % (((^(v15))) | 1))) | 1))}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    if (true || (false || ((v2 >= 1.5) || false))) {
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := GenS1{F0: 42, F1: genFn1()}
            fmt.Println(fi17, fi17, v15)
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 string = ((("\n\t" + "hello") + "\n\t") + "hello")
                v21 := GenS2{F0: 396}
                if true {
                    fmt.Println(v3, v20, fi19)
                    var v22 uint64 = 42
                    var v23 uint16 = 945
                    _ = v22
                    _ = v23
                }
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi17
            _ = v18
        }
        fmt.Println(v2, v15)
        for fi24 := 0; fi24 < 16; fi24++ {
            if (-128 == 127) {
                v2 -= v2
            }
            var v25 uint16 = 32767
            _ = fi24
            _ = v25
        }
    }
    delete(v4, (-(v15)))
    for fi26 := 0; fi26 < 15; fi26++ {
        v2 = v2
        if len(v16) > 0 {
            v16[0] = fi26
        }
        v27 := GenS3{F0: v1, F1: 500}
        _ = fi26
        _ = v27
    }
    v3 = (v3 >> ((2) & 63))
    for fi28 := 0; fi28 < 19; fi28++ {
        var v29 uint = 531
        if false {
            fmt.Println(v2)
        } else {
            if true {
                for fi30 := 0; fi30 < 2; fi30++ {
                    v3 = v3
                    _ = fi30
                }
                v31 := GenS2{F0: -1}
                if false {
                    v32 := GenS4{F0: v1}
                    for fi33 := 0; fi33 < 8; fi33++ {
                        var v34 uint = 2
                        v4[616] = (("hello" + "\n\t") < "\n\t")
                        if ((0.0 / v2) >= (float64((-2.25 + 0.0)) - 1.5)) {
                            var v35 bool = (42 != 2)
                            _ = v35
                        } else {
                            switch (("\n\t" + "foo bar") + "\n\t") {
                            case "hello":
                                fmt.Println(v3)
                            case "\n\t":
                                var v36 uint16 = 255
                                _ = v36
                            case "foo bar":
                                v3 -= v3
                            }
                        }
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                } else {
                    var v37 int64 = -320
                    delete(v4, (v15 | int(v29)))
                    for fi38 := 0; fi38 < 6; fi38++ {
                        v39 := GenS1{F0: 42, F1: 4294967295}
                        _ = fi38
                        _ = v39
                    }
                    _ = v37
                }
                _ = v31
            }
            fmt.Println(v29, v15)
        }
        var v40 int8 = 127
        _ = fi28
        _ = v29
        _ = v40
    }
    _ = len(v4)
    v2 -= ((3.14159 * v2) * -2.25)
    for fi41 := 0; fi41 < 16; fi41++ {
        switch "hello" {
        case "":
            var v42 int32 = 0
            _ = v42
        case "\n\t":
            v3 += (^(v3))
        }
        v3 = 0
        if true {
            var v43 int32 = -1
            _ = v43
        } else {
            for fi44 := 0; fi44 < 9; fi44++ {
                if true {
                    var v45 float32 = (0.0 + -2.25)
                    _ = v45
                }
                _ = fi44
            }
            fmt.Println(fi41, v2, v3)
            fmt.Println(fi41, v15)
        }
        _ = fi41
    }
    v2 += v2
    var v46 uint64 = 4294967295
    var v47 float32 = 1e10
    fmt.Println(v2, v3, v47)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
}
