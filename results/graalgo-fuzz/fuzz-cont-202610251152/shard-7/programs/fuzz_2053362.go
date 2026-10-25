package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int16; F2 uint64 }
func genFn0(p0 int8, p1 string) int16 {
    var v4 int64 = 960
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 255}
        var v7 uint64 = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: 587, F1: 32767}
    fmt.Println(v4, p0)
    v4 += (-(v4))
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return 42
}
func genFn1(p0 int16) int32 {
    v9 := []int{}
    switch "" {
    case "\n\t":
        v10 := GenS3{F0: 652, F1: 1}
        _ = v10
    case "a":
        _ = len(v9)
    case "foo bar":
        if (42 == 76) {
            for fi11 := 0; fi11 < 7; fi11++ {
                switch ("foo bar" + "hello") {
                case "":
                    v12 := GenS1{F0: 2}
                    _ = v12
                case "foo bar":
                    if false {
                        if false {
                            v9 = append(v9, (fi11 ^ -79))
                            p0 = p0
                        } else {
                            v13 := GenS4{F0: 816}
                            _ = v13
                        }
                    }
                }
                p0 += 1
                _ = fi11
            }
            var v14 uint8 = 255
            _ = v14
        } else {
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 uint64 = 620
                for fi17 := 0; fi17 < 17; fi17++ {
                    p0 = p0
                    _ = fi17
                }
                if (0.0 <= -2.25) {
                    var v18 int64 = -1
                    _ = v18
                } else {
                    v19 := GenS1{F0: 0}
                    for fi20 := 0; fi20 < 6; fi20++ {
                        p0 = p0
                        v21 := GenS1{F0: 2}
                        if true {
                            var v22 int32 = 1
                            fmt.Println(v22, fi20, fi15)
                            _ = v22
                        }
                        _ = fi20
                        _ = v21
                    }
                    fmt.Println(fi15, v16)
                    _ = v19
                }
                _ = fi15
                _ = v16
            }
            if ((0.0 != 3.14159) && false) {
                var v23 uint32 = 2
                var v24 int16 = (-(p0))
                var v25 uint32 = 42
                _ = v23
                _ = v24
                _ = v25
            } else {
                for fi26 := 0; fi26 < 13; fi26++ {
                    p0 = p0
                    if true {
                        var v27 uint16 = 0
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 14; fi28++ {
                            v29 := GenS4{F0: 2}
                            switch "foo bar" {
                            case "a":
                                v30 := GenS5{F0: 255, F1: (p0 | genFn0(-50, ("a" + ("hello" + "foo bar")))), F2: 1}
                                _ = v30
                            }
                            _ = fi28
                            _ = v29
                        }
                        switch ("\n\t" + "hello") {
                        case "":
                            switch "foo bar" {
                            case "":
                                switch ("foo bar" + "\n\t") {
                                case "a":
                                    for fi31 := 0; fi31 < 18; fi31++ {
                                        p0 -= p0
                                        var v32 int32 = 323
                                        var v33 float64 = 1.5
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                    }
                                case "\n\t":
                                    for fi34 := 0; fi34 < 13; fi34++ {
                                        fmt.Println(fi26, fi26, fi26)
                                        _ = fi34
                                    }
                                }
                            case "foo bar":
                                if true {
                                    var v35 string = "\n\t"
                                    switch fi26 {
                                    case 32767:
                                        _ = len(v9)
                                    case -2147483648:
                                        v35 = ((v35 + v35) + v35)
                                    case 4294967295:
                                        p0 = 42
                                    }
                                    _ = v35
                                } else {
                                    v9 = append(v9, 114)
                                    v36 := GenS2{F0: 1, F1: 100}
                                    var v37 uint = 367
                                    _ = v36
                                    _ = v37
                                }
                            case "a":
                                for fi38 := 0; fi38 < 6; fi38++ {
                                    var v39 int64 = -1
                                    var v40 uint32 = 803
                                    p0 = (-(p0))
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                            }
                        case "a":
                            if false {
                                p0 += genFn0(-48, ("hello" + (("hello" + "\n\t") + "hello")))
                            }
                        }
                        p0 = genFn0(100, "a")
                    }
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = fi26
                }
            }
        }
    }
    _ = p0
    _ = v9
    return -144
}

func main() {
    var v1 uint = 4294967295
    var v2 uint8 = 1
    v3 := make(map[string]string)
    for fi41 := 0; fi41 < 15; fi41++ {
        v1 = (^(v1))
        v42 := []int{fi41, -1, (fi41 * fi41), fi41}
        _ = fi41
        _ = v42
    }
    switch 32767 {
    case -1:
        v2 += 1
    case 32767:
        var v43 uint64 = 854
        _ = v43
    case 9223372036854775807:
        v2 = v2
    }
    fmt.Println(v2, v2)
    v44 := []int{4294967295, 4294967295, 32767}
    var v45 string = (("foo bar" + (("a" + (("\n\t" + "\n\t") + "foo bar")) + "\n\t")) + (("a" + "\n\t") + ("hello" + "\n\t")))
    var v46 int64 = -782
    v47 := GenS1{F0: v2}
    if (683 >= 2147483647) {
        v2 = v2
        fmt.Println(v1, v1, v1)
    }
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    var v48 uint32 = 2147483648
    var v49 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
