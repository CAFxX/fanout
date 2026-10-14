package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int32; F2 int32 }
type GenS3 struct { F0 int8; F1 bool; F2 int8 }
func genFn0(p0 int8, p1 uint32) int16 {
    var v7 int32 = 1
    fmt.Println(p0, p0)
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 0}
        switch (^(fi8)) {
        case 127:
            switch (("hello" + "a") + ("" + "hello")) {
            case "foo bar":
                fmt.Println(p0, p1, p1)
            case "hello":
                v10 := GenS1{F0: 1}
                _ = v10
            default:
                var v11 int32 = v7
                _ = v11
            }
        default:
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 uint32 = (p1 - (^(p1)))
                if false {
                    fmt.Println(v7, p1)
                    var v14 uint32 = (v13 & (p1 >> ((366) & 63)))
                    _ = v14
                }
                switch -594 {
                case 42:
                    v15 := []int{fi12, (fi12 + fi12)}
                    _ = v15
                case 127:
                    if (292 > 42) {
                        fmt.Println(v13, v7, v7)
                        v16 := GenS1{F0: 18446744073709551615}
                        _ = v16
                    }
                }
                _ = fi12
                _ = v13
            }
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := GenS1{F0: 18446744073709551615}
            v7 = (v7 - int32(p0))
            v19 := GenS1{F0: 1}
            v20 := []int{fi8, fi8}
            switch ("\n\t" + "hello") {
            case "":
                v21 := GenS2{F0: p0, F1: -1, F2: v7}
                _ = v21
            case "hello":
                v7 -= (-(v7))
            }
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi8
        _ = v9
    }
    switch "a" {
    case "hello":
        v22 := []int{614, -690, 2147483647, 1}
        _ = v22
    default:
        switch "\n\t" {
        case "a":
            p0 -= (p0 ^ 1)
        default:
            fmt.Println(p0, p0)
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1() uint {
    var v23 string = ""
    v23 = v23
    if true {
        fmt.Println(v23)
        if false {
            fmt.Println(v23)
        }
        v23 = "\n\t"
    }
    if true {
        v23 = v23
    }
    _ = v23
    return 255
}

func main() {
    v1 := []string{"\n\t", ("" + "")}
    var v2 int8 = 127
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 127
    v6 := make(map[string]bool)
    switch 604 {
    case 127:
        v5 = v5
    case 4294967295:
        var v24 int = -249
        _ = v24
    case -1:
        v25 := GenS1{F0: 206}
        _ = v25
    default:
        v26 := GenS3{F0: v2, F1: false, F2: (v2 & v2)}
        _ = v26
    }
    switch ("hello" + "") {
    case "":
        var v27 float32 = float32(1.5)
        _ = v27
    }
    var v28 uint8 = v5
    v2 -= 42
    var v29 int64 = -1
    var v30 bool = (1 <= 42)
    for fi31 := 0; fi31 < 7; fi31++ {
        var v32 int32 = -2147483648
        for fi33 := 0; fi33 < 6; fi33++ {
            var v34 int32 = (v32 + (v32 | v32))
            _ = fi33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    fmt.Println(v2)
    v35 := v6[(("" + "hello") + (("\n\t" + "hello") + ""))]
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 uint = 255
        v28 = v5
        v38 := GenS3{F0: (v2 / ((-84) | 1)), F1: (-32768 < 42), F2: int8(v29)}
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 uint = 18446744073709551615
    for fi40 := 0; fi40 < 13; fi40++ {
        var v41 string = "a"
        v5 += 2
        _ = fi40
        _ = v41
    }
    delete(v4, 42)
    var v42 int = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
