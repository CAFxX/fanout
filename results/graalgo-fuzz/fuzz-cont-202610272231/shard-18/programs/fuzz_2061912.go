package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint; F1 uint32; F2 int8 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 int64, p1 string, p2 uint32) int64 {
    var v5 uint16 = 418
    if false {
        var v6 int = -9223372036854775808
        p0 += (p0 | -902)
        _ = v6
    } else {
        if false {
            var v7 float32 = -2.25
            _ = v7
        } else {
            var v8 int32 = 42
            var v9 float32 = ((-2.25 * (1.5 * (0.0 * 1.5))) + 1.5)
            _ = v8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := []int{-20, fi10}
        _ = fi10
        _ = v11
    }
    v12 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v5
    return (-(p0))
}
func genFn1(p0 int32, p1 uint32) bool {
    var v13 uint32 = p1
    if ("\n\t" > ("" + ("" + "foo bar"))) {
        var v14 int32 = p0
        switch -533 {
        case 757:
            v15 := GenS1{F0: 32767}
            _ = v15
        case -1:
            var v16 int32 = 2147483647
            _ = v16
        case 4294967295:
            fmt.Println(v14)
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            v18 := []int{fi17, (-(fi17)), (fi17 + fi17)}
            v14 = p0
            _ = fi17
            _ = v18
        }
        _ = v14
    } else {
        switch "" {
        case "a":
            v13 = 2147483647
        case "\n\t":
            fmt.Println(p1, p1, p0)
        case "hello":
            p0 += (p0 + (^(p0)))
        }
        fmt.Println(p0, v13, v13)
        v19 := GenS1{F0: 45}
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v13
    return false
}

func main() {
    var v1 string = ("\n\t" + ((("foo bar" + ("\n\t" + "\n\t")) + "") + ("hello" + ("hello" + "\n\t"))))
    var v2 int8 = 7
    var v3 uint32 = 578
    v4 := []string{("" + "foo bar"), v1}
    var v20 string = v1
    switch -177 {
    case 42:
        var v21 uint64 = 42
        _ = v21
    case 0:
        v4 = append(v4, v20)
    case 127:
        v4 = append(v4, v20)
    }
    v3 = 4294967295
    if (int(v3) == 2147483647) {
        for fi22 := 0; fi22 < 4; fi22++ {
            if ((!true) || ((v1 + (v20 + (v20 + v20))) > "foo bar")) {
                if (((true && (85 != (v2 >> ((0) & 63)))) || false) && false) {
                    if len(v4) > 0 {
                        v4[0] = (((v20 + v20) + v20) + v1)
                    }
                    var v23 string = ""
                    var v24 int = (^(fi22))
                    _ = v23
                    _ = v24
                }
            }
            if (1 > 1) {
                _ = len(v4)
                var v25 int = fi22
                _ = v25
            } else {
                switch v20 {
                case "\n\t":
                    v4 = append(v4, "a")
                default:
                    switch fi22 {
                    case -234:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        var v26 int64 = 0
                        _ = v26
                    }
                }
                v3 += (v3 * (v3 ^ v3))
                v27 := GenS2{F0: ""}
                _ = v27
            }
            var v28 int = (fi22 << ((690) & 63))
            _ = fi22
            _ = v28
        }
        v1 = "hello"
        v29 := GenS2{F0: ((v20 + ((v20 + v1) + "foo bar")) + (((((v1 + (((v20 + (v20 + (v20 + v20))) + v1) + "a")) + "") + v1) + v20) + ("" + ("hello" + v20))))}
        _ = v29
    } else {
        var v30 uint32 = 1
        v30 = v3
        _ = v30
    }
    if (-2.25 != 0.1) {
        v31 := GenS1{F0: 0}
        fmt.Println(v1, v2, v1)
        for fi32 := 0; fi32 < 4; fi32++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch (v1 + v20) {
            case "":
                switch (v20 + (v1 + v1)) {
                case "a":
                    _ = len(v4)
                case "":
                    v2 = (v2 >> ((255) & 63))
                default:
                    for fi33 := 0; fi33 < 16; fi33++ {
                        v34 := GenS3{F0: 18446744073709551615, F1: 416}
                        v1 = ("" + ("a" + "foo bar"))
                        _ = fi33
                        _ = v34
                    }
                }
            case "a":
                if (255 > 110) {
                    var v35 bool = true
                    v2 += 100
                    _ = v35
                } else {
                    var v36 int32 = 819
                    if true {
                        var v37 uint32 = 0
                        _ = v37
                    }
                    _ = v36
                }
            case "foo bar":
                fmt.Println(v3)
            }
            _ = fi32
        }
        _ = v31
    } else {
        switch 32767 {
        case -1:
            for fi38 := 0; fi38 < 14; fi38++ {
                var v39 int = fi38
                var v40 int16 = -32768
                _ = fi38
                _ = v39
                _ = v40
            }
        }
        v41 := GenS4{F0: 42, F1: v3, F2: (v2 + v2)}
        if true {
            v42 := GenS5{F0: -32768, F1: false}
            v43 := GenS6{F0: 0}
            v1 = v1
            _ = v42
            _ = v43
        } else {
            fmt.Println(v1)
        }
        _ = v41
    }
    switch 1 {
    case 32767:
        var v44 int16 = 1000
        _ = v44
    }
    v4 = append(v4, "")
    var v45 uint8 = 2
    fmt.Println(v3)
    v46 := GenS7{F0: (v3 * 489)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v45
    _ = v46
}
