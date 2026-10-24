package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint32, p1 int8, p2 uint64) uint32 {
    p2 += 498
    var v7 int64 = -601
    v8 := GenS1{F0: (p2 & 4294967295)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p0 & (^(p0)))
}

func main() {
    var v1 uint = 291
    var v2 string = "hello"
    var v3 uint32 = 728
    v4 := make(map[string]string)
    var v5 string = ("" + ((v2 + v2) + v2))
    var v6 float64 = 1e10
    var v9 uint64 = 42
    v10 := GenS1{F0: v9}
    v11 := GenS1{F0: v9}
    v4[(v2 + v5)] = (v5 + (v5 + v5))
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint8 = 64
        _ = fi12
        _ = v13
    }
    v2 = (v2 + (v5 + v5))
    var v14 int64 = 1
    v15 := GenS2{F0: 100}
    var v16 int = -1
    var v17 float32 = float32(v6)
    v5 = v5
    for fi18 := 0; fi18 < 16; fi18++ {
        if false {
            for fi19 := 0; fi19 < 3; fi19++ {
                if ((-1 > -799) && ((((v2 + ((v2 + v2) + ("hello" + "hello"))) + ("\n\t" + v5)) + "hello") == v2)) {
                    v20 := GenS3{F0: (v3 >> ((1) & 63)), F1: (v3 & (v3 - (v3 + (v3 - v3)))), F2: (^(v3))}
                    _ = v20
                } else {
                    var v21 int8 = 127
                    var v22 uint32 = 951
                    var v23 int = fi18
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = fi19
            }
        } else {
            v14 += v14
            var v24 string = ((v2 + v5) + ("hello" + (((((v2 + "\n\t") + (("foo bar" + ("\n\t" + "foo bar")) + "")) + (("hello" + "a") + ("\n\t" + v2))) + "") + (v5 + (((v2 + ((v5 + v5) + ("\n\t" + (v2 + v2)))) + (((v5 + (v5 + v5)) + "") + "foo bar")) + "foo bar")))))
            _ = v24
        }
        if true {
            if true {
                switch v2 {
                case "foo bar":
                    v25 := GenS4{F0: (v9 % (((v9 ^ v9)) | 1)), F1: (v14 >> ((v1) & 63))}
                    _ = v25
                default:
                    v26 := v4[((("hello" + v5) + v2) + v5)]
                    _ = v26
                }
                v27 := GenS1{F0: (^(v9))}
                if (4294967295 == 0) {
                    if true {
                        for fi28 := 0; fi28 < 18; fi28++ {
                            delete(v4, ((v2 + v5) + (v5 + v5)))
                            v6 = v6
                            _ = fi28
                        }
                        v4[((v5 + "") + (v5 + v2))] = v2
                        var v29 float32 = float32(v6)
                        _ = v29
                    } else {
                        if true {
                            var v30 int = fi18
                            _ = v30
                        }
                        v31 := GenS5{F0: 505}
                        v16 = fi18
                        _ = v31
                    }
                } else {
                    if false {
                        var v32 bool = false
                        fmt.Println(v2, fi18, v14)
                        _ = v32
                    }
                }
                _ = v27
            }
        }
        _ = fi18
    }
    v33 := GenS3{F0: genFn0(v3, -128, v9), F1: (v3 - (v3 ^ v3)), F2: 199}
    var v34 int8 = 121
    var v35 uint32 = 685
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
