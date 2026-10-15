package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int64; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 16
    var v3 bool = true
    var v4 uint16 = 0
    var v5 uint32 = 2
    v6 := GenS1{F0: -1}
    switch "\n\t" {
    case "a":
        v4 -= v4
    case "foo bar":
        v7 := GenS1{F0: -2147483648}
        _ = v7
    case "\n\t":
        if (!(1e10 == 1e10)) {
            switch ("foo bar" + ("foo bar" + ((("\n\t" + "\n\t") + "\n\t") + "foo bar"))) {
            case "":
                v5 = 1
            case "\n\t":
                v8 := GenS2{F0: -1, F1: (v2 - v2)}
                _ = v8
            case "a":
                v9 := GenS1{F0: -918}
                _ = v9
            }
            v10 := GenS3{F0: -1, F1: 933, F2: 42}
            _ = v10
        } else {
            v5 = v5
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 int16 = -1
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 uint = 65535
                    v12 = int16(fi11)
                    _ = fi13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
        }
    default:
        var v15 uint8 = 42
        _ = v15
    }
    v5 -= (v5 - v5)
    fmt.Println(v3, v3)
    fmt.Println(v2)
    var v16 uint16 = v4
    var v17 uint64 = 42
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 6; fi19++ {
            v20 := GenS1{F0: -2147483648}
            v5 = 191
            _ = fi19
            _ = v20
        }
        fmt.Println(v16)
        for fi21 := 0; fi21 < 5; fi21++ {
            if (!v3) {
                switch fi21 {
                case 938:
                    _ = len(v1)
                default:
                    if true {
                        var v22 string = ("a" + ("hello" + (("" + (("\n\t" + "\n\t") + ("\n\t" + "a"))) + "hello")))
                        v23 := GenS3{F0: 42, F1: 247, F2: int64(v4)}
                        v24 := v1[42]
                        _ = v22
                        _ = v23
                        _ = v24
                    } else {
                        var v25 int32 = 880
                        v17 += (^(v17))
                        v26 := []int{(fi18 - 32767), (fi21 << ((910) & 63)), fi18, 42}
                        _ = len(v1)
                        if v3 {
                            var v27 int = 9223372036854775807
                            v17 = v17
                            fmt.Println(v3, v5, v17)
                            _ = v27
                        } else {
                            var v28 uint = 73
                            var v29 int32 = (v25 % ((v25) | 1))
                            _ = v28
                            _ = v29
                        }
                        _ = v25
                        _ = v26
                    }
                    delete(v1, 128)
                    fmt.Println(fi18, v16)
                }
                if ((!(!(!(!false)))) && ((v3 && (-1 >= 1000)) && (v16 != v16))) {
                    v30 := GenS2{F0: 42, F1: 127}
                    var v31 uint16 = (v16 << ((255) & 63))
                    v17 = (v17 * 42)
                    _ = v30
                    _ = v31
                }
                fmt.Println(v3)
            } else {
                var v32 float64 = 0.0
                var v33 int64 = -108
                _ = v32
                _ = v33
            }
            _ = fi21
        }
        _ = fi18
    }
    v2 = (v2 % ((-32) | 1))
    v3 = (v5 > 278)
    var v34 bool = (1 > 979)
    if v34 {
        for fi35 := 0; fi35 < 19; fi35++ {
            for fi36 := 0; fi36 < 19; fi36++ {
                v37 := v1[(fi35 ^ -371)]
                fmt.Println(v17, v4)
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
        if v3 {
            v38 := GenS2{F0: -1, F1: (-(v2))}
            _ = v38
        }
    }
    var v39 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
