package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 int32; F1 int8 }
func genFn0(p0 int8, p1 uint, p2 uint64) int8 {
    var v5 uint8 = 1
    var v6 uint = 4294967295
    if true {
        fmt.Println(v6)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p0 >> (((^(p1))) & 63))
}
func genFn1(p0 uint64, p1 int, p2 uint) bool {
    if (false || true) {
        if false {
            for fi7 := 0; fi7 < 14; fi7++ {
                p2 += 255
                fmt.Println(p2, p2, p0)
                v8 := GenS1{F0: (p2 / ((p2) | 1)), F1: 439}
                _ = fi7
                _ = v8
            }
            v9 := []int{(p1 / ((p1) | 1))}
            _ = len(v9)
            _ = v9
        } else {
            for fi10 := 0; fi10 < 14; fi10++ {
                p0 = (p0 >> ((p2) & 63))
                _ = fi10
            }
            var v11 uint8 = 180
            _ = v11
        }
        fmt.Println(p2, p1, p2)
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 string = (("foo bar" + "") + "")
        _ = fi12
        _ = v13
    }
    if (false && true) {
        p1 -= p1
    }
    var v14 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (!(p1 <= p1))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    v15 := GenS1{F0: 4294967295, F1: 42}
    var v16 int64 = -1
    v17 := v3["hello"]
    if (float32(-2.25) < 1e10) {
        v18 := GenS1{F0: 168, F1: 2}
        _ = v18
    } else {
        switch -2147483648 {
        case 32768:
            fmt.Println(v16)
        }
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := []int{4294967295, fi19, (fi19 + (fi19 + -1))}
            for fi21 := 0; fi21 < 3; fi21++ {
                if v17 {
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    var v22 uint64 = 87
                    _ = v22
                } else {
                    var v23 bool = (127 <= 1)
                    _ = len(v20)
                    _ = v23
                }
                var v24 int64 = v16
                fmt.Println(v24)
                _ = fi21
                _ = v24
            }
            fmt.Println(v17, v16, fi19)
            _ = fi19
            _ = v20
        }
        v17 = true
    }
    var v25 float64 = (1.5 + 0.1)
    switch (("" + (("hello" + "a") + "foo bar")) + ((((("\n\t" + "a") + "hello") + "a") + "hello") + "foo bar")) {
    case "a":
        fmt.Println(v17, v16, v25)
    case "foo bar":
        switch -441 {
        case 4294967295:
            v16 = (v16 - v16)
        case 32767:
            switch -1 {
            case 527:
                switch ("foo bar" + ("hello" + ((("" + ((("foo bar" + ("foo bar" + (("a" + "") + ((("\n\t" + "") + "a") + ("\n\t" + (("foo bar" + "hello") + (("foo bar" + "\n\t") + "a"))))))) + (((("\n\t" + ("\n\t" + ((("\n\t" + "a") + "\n\t") + "a"))) + ("hello" + "a")) + ((("foo bar" + "foo bar") + ("hello" + ("foo bar" + "a"))) + "a")) + "foo bar")) + "foo bar")) + "\n\t") + ""))) {
                case "":
                    delete(v3, "a")
                default:
                    v25 = ((v25 - v25) - 0.0)
                }
            default:
                var v26 uint = 619
                _ = v26
            }
        }
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        var v28 int32 = -446
        v29 := GenS1{F0: 255, F1: 1}
        switch (((((("\n\t" + "hello") + ("foo bar" + "\n\t")) + "") + ("\n\t" + ("hello" + "\n\t"))) + ("hello" + "\n\t")) + "a") {
        case "foo bar":
            for fi30 := 0; fi30 < 17; fi30++ {
                v31 := GenS2{F0: (^(v16))}
                fmt.Println(fi27, v28)
                v17 = false
                _ = fi30
                _ = v31
            }
        default:
            if v17 {
                var v32 float32 = 1.5
                var v33 float32 = 3.14159
                _ = v32
                _ = v33
            } else {
                v28 = 1
                switch "\n\t" {
                case "\n\t":
                    switch "foo bar" {
                    case "a":
                        v25 = v25
                    default:
                        if v17 {
                            v16 = (v16 / (((v16 << ((42) & 63))) | 1))
                            var v34 float32 = 0.0
                            _ = v34
                        } else {
                            var v35 uint = 2
                            _ = v35
                        }
                    }
                case "foo bar":
                    v25 -= 1e10
                case "":
                    v36 := GenS3{F0: 18446744073709551615}
                    _ = v36
                }
            }
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    if (-563 != 2147483647) {
        var v37 int64 = -9223372036854775808
        switch -853 {
        case -9223372036854775808:
            _ = len(v1)
        case -1:
            switch 127 {
            case 0:
                fmt.Println(v16, v25)
            default:
                var v38 int32 = -1
                _ = v38
            }
        case 42:
            fmt.Println(v17, v16)
        default:
            switch "" {
            case "\n\t":
                var v39 string = ("\n\t" + ("foo bar" + ("" + "foo bar")))
                _ = v39
            case "a":
                switch "foo bar" {
                case "a":
                    var v40 string = ""
                    _ = v40
                }
            default:
                var v41 int64 = v37
                _ = v41
            }
        }
        var v42 float64 = v25
        _ = v37
        _ = v42
    } else {
        fmt.Println(v25, v17)
    }
    v16 = (v16 >> ((928) & 63))
    switch "foo bar" {
    case "foo bar":
        v16 -= -1
    case "a":
        v43 := GenS4{F0: 0, F1: 42}
        _ = v43
    case "hello":
        if v17 {
            v44 := GenS1{F0: 65535, F1: 42}
            v45 := GenS2{F0: (v16 >> ((255) & 63))}
            _ = v44
            _ = v45
        } else {
            var v46 uint64 = 18446744073709551615
            v47 := GenS5{F0: 325, F1: 0}
            _ = v46
            _ = v47
        }
    }
    fmt.Println(v16, v16)
    var v48 uint = 18446744073709551615
    v49 := GenS6{F0: -369, F1: 100}
    v48 += 2
    v17 = false
    fmt.Println(v48)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v48
    _ = v49
}
