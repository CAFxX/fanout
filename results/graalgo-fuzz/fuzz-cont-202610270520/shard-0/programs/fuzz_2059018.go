package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int; F2 int64 }
type GenS4 struct { F0 bool; F1 bool; F2 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 uint; F2 string }
type GenS7 struct { F0 int; F1 bool; F2 string }
type GenS8 struct { F0 int32; F1 int64; F2 int32 }
func genFn0(p0 int32, p1 bool) string {
    if (!true) {
        v6 := []int{}
        if p1 {
            v7 := GenS1{F0: 2147483647, F1: (("a" + "\n\t") + (("hello" + "a") + ("\n\t" + "foo bar")))}
            switch ("foo bar" + ("a" + "\n\t")) {
            case "foo bar":
                var v8 bool = (true || false)
                _ = v8
            case "hello":
                var v9 bool = (-9223372036854775808 > 2147483647)
                _ = v9
            case "\n\t":
                _ = len(v6)
            }
            switch ("\n\t" + ("a" + ("foo bar" + "a"))) {
            case "hello":
                _ = len(v6)
            case "a":
                p0 += (p0 / (((p0 * p0)) | 1))
            case "foo bar":
                p1 = (!((1 > 717) || p1))
            default:
                p1 = (-9223372036854775808 < 464)
            }
            _ = v7
        } else {
            for fi10 := 0; fi10 < 4; fi10++ {
                p1 = ((377 <= 32767) || p1)
                _ = fi10
            }
            for fi11 := 0; fi11 < 18; fi11++ {
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := GenS1{F0: 223, F1: "foo bar"}
                    if (p1 || p1) {
                        p0 = (p0 >> ((615) & 63))
                    }
                    var v14 int32 = p0
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                if true {
                    p0 += (p0 - (p0 ^ p0))
                    var v15 int = fi11
                    v16 := GenS2{F0: 255}
                    _ = v15
                    _ = v16
                }
                _ = fi11
            }
        }
        var v17 int64 = -918
        _ = v17
        _ = v6
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := []int{fi19}
            switch "hello" {
            case "hello":
                v21 := GenS3{F0: (-925 <= p0), F1: (fi18 & fi18), F2: 1}
                p0 = (p0 - (p0 * p0))
                fmt.Println(fi19)
                _ = v21
            }
            v22 := GenS4{F0: ((p1 && (-9223372036854775808 < 9223372036854775807)) || p1), F1: true, F2: 0}
            _ = fi19
            _ = v20
            _ = v22
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ("" + "a")
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]bool)
    v5 := []string{"\n\t", (v2 + (v2 + "foo bar"))}
    var v23 string = "hello"
    var v24 int32 = 370
    var v25 uint32 = (v3 << ((65535) & 63))
    v26 := GenS3{F0: false, F1: 32768, F2: -9223372036854775808}
    v3 -= (v25 | (v25 & v25))
    v27 := GenS5{F0: 2}
    var v28 uint16 = 32767
    switch 757 {
    case 9223372036854775807:
        for fi29 := 0; fi29 < 5; fi29++ {
            switch v23 {
            case "":
                switch v2 {
                case "\n\t":
                    v30 := GenS6{F0: 100, F1: 255, F2: v23}
                    _ = v30
                case "hello":
                    v3 -= 2147483648
                }
            }
            _ = fi29
        }
    case 128:
        var v31 string = v23
        _ = v31
    }
    v2 = v2
    fmt.Println(v23)
    for fi32 := 0; fi32 < 2; fi32++ {
        v33 := GenS7{F0: (fi32 ^ (fi32 / ((fi32) | 1))), F1: true, F2: "\n\t"}
        var v34 int16 = int16(v24)
        delete(v1, (v2 + v23))
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 int32 = v24
    for fi36 := 0; fi36 < 2; fi36++ {
        var v37 int16 = -32768
        var v38 uint32 = 2147483648
        v39 := GenS8{F0: v35, F1: 42, F2: 1}
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    if len(v5) > 0 {
        v5[0] = v2
    }
    if len(v5) > 0 {
        v5[0] = (("\n\t" + ((v2 + "foo bar") + "hello")) + v23)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
