package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0() int {
    var v5 float32 = 1.5
    v6 := GenS1{F0: -63, F1: 18446744073709551615, F2: 200}
    v5 = (v5 * (v5 - (0.1 / v5)))
    _ = v5
    _ = v6
    return -990
}
func genFn1(p0 int32, p1 int32, p2 string) int32 {
    if false {
        fmt.Println(p2, p2, p1)
        for fi7 := 0; fi7 < 19; fi7++ {
            if (false && false) {
                switch (fi7 * (fi7 - genFn0())) {
                case -9223372036854775808:
                    var v8 uint64 = 740
                    _ = v8
                case -112:
                    if false {
                        var v9 uint8 = 198
                        _ = v9
                    }
                case 127:
                    var v10 uint = 18446744073709551615
                    _ = v10
                }
            } else {
                fmt.Println(p1, p2, p2)
                p0 -= 0
            }
            var v11 uint = 576
            p0 = -2147483648
            _ = fi7
            _ = v11
        }
        fmt.Println(p0)
    }
    v12 := []int{32767, 128}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v13 int8 = 63
    v14 := GenS1{F0: genFn0(), F1: 249, F2: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 782
    v3 := make(map[string]string)
    var v4 bool = true
    v4 = v4
    var v15 string = ""
    var v16 string = v15
    var v17 bool = (v4 || v4)
    v18 := GenS1{F0: 937, F1: 61, F2: 128}
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 string = ""
        if (!true) {
            v21 := GenS1{F0: fi19, F1: 18446744073709551615, F2: uint8(v2)}
            if v17 {
                fmt.Println(v15, v2, v16)
                for fi22 := 0; fi22 < 16; fi22++ {
                    if v4 {
                        var v23 string = (v16 + "hello")
                        switch v20 {
                        case "a":
                            v24 := GenS1{F0: genFn0(), F1: 0, F2: 216}
                            _ = v24
                        }
                        switch fi22 {
                        case -545:
                            v20 = ((v15 + (v23 + v16)) + "a")
                        case 1:
                            v25 := []int{genFn0(), (-(fi22)), fi19}
                            _ = v25
                        case -1:
                            var v26 int32 = -1
                            _ = v26
                        }
                        var v27 float32 = 0.1
                        v28 := v3[("\n\t" + (("a" + "") + ""))]
                        v23 = "a"
                        switch 111 {
                        case 32768:
                            if (42 >= 0) {
                                if (-1 < 9223372036854775807) {
                                    v23 = v15
                                    if false {
                                        v3[v28] = (v20 + "")
                                        fmt.Println(v27, fi22, fi19)
                                        if false {
                                            v16 = v23
                                            var v29 uint16 = 1
                                            _ = v29
                                        }
                                    }
                                    var v30 float32 = ((v27 / v27) * 1.5)
                                    _ = v30
                                } else {
                                    var v31 string = v16
                                    v2 = 402
                                    var v32 int64 = 1
                                    _ = v31
                                    _ = v32
                                }
                            }
                        case 42:
                            v3[((v16 + "foo bar") + "foo bar")] = v15
                        }
                        _ = v23
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(fi19, fi19)
                    v33 := v3["hello"]
                    _ = fi22
                    _ = v33
                }
                if ((0.0 / 1e10) == -2.25) {
                    var v34 int = (fi19 * (fi19 + fi19))
                    _ = v34
                }
            } else {
                v17 = v17
                v35 := GenS2{F0: (32767 < 1000), F1: 128}
                _ = v35
            }
            var v36 int = fi19
            _ = v21
            _ = v36
        }
        _ = fi19
        _ = v20
    }
    fmt.Println(v15)
    _ = len(v1)
    for fi37 := 0; fi37 < 17; fi37++ {
        switch (v15 + (("hello" + ((v16 + v16) + "\n\t")) + v16)) {
        case "":
            var v38 uint16 = 2
            _ = v38
        }
        v39 := GenS1{F0: genFn0(), F1: 4294967295, F2: 2}
        _ = fi37
        _ = v39
    }
    var v40 string = ("a" + (((v16 + "foo bar") + v15) + ""))
    if (!false) {
        v16 = (v16 + v15)
    }
    for fi41 := 0; fi41 < 19; fi41++ {
        v42 := GenS2{F0: (!false), F1: 42}
        _ = fi41
        _ = v42
    }
    v4 = v4
    for fi43 := 0; fi43 < 8; fi43++ {
        v44 := GenS3{F0: 255, F1: fi43}
        v16 = v16
        _ = fi43
        _ = v44
    }
    v3[v15] = v40
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}
